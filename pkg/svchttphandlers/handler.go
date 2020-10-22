package svchttphandlers

import (
	"context"
	"encoding/json"

	"github.com/go-kit/kit/log"

	"net/http"
	"strings"

	httptransport "github.com/go-kit/kit/transport/http"
)

type contextKey int

const (
	// ContextKeyPermissionHeader sets the context key for permission header `X-Viki-App-Permissions`
	ContextKeyPermissionHeader contextKey = iota
)

// DefaultServerOptions returns the default server options used by transport layers.
func DefaultServerOptions(logger log.Logger) []httptransport.ServerOption {
	return []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
		httptransport.ServerBefore(extractPermission),
	}
}

// EncodeStatusCreatedResponse encodes response struct to return for clients for 201 Created request.
func EncodeStatusCreatedResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.Error() != nil {
		encodeError(ctx, e.Error(), w)
		return e.Error()
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(response)
}

// EncodeStatusOKResponse encodes response struct to return for clients for 200 OK request.
func EncodeStatusOKResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.Error() != nil {
		encodeError(ctx, e.Error(), w)
		return e.Error()
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	errResp := getErrorResponse(err)
	w.WriteHeader(codeFrom(errResp))
	json.NewEncoder(w).Encode(errResp)
}

type errorer interface {
	Error() error
}

func codeFrom(errResp ErrorResponse) int {
	switch errResp.VCode {
	case ErrNotFound:
		return http.StatusNotFound

	case ErrInternalCode:
		return http.StatusInternalServerError

	case ErrBadRequestCode:
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}

func extractPermission(ctx context.Context, req *http.Request) context.Context {
	permissions := strings.Split(req.Header.Get(`X-Viki-App-Permissions`), ",")
	availablePermissions := make(map[string]bool)
	for _, permission := range permissions {
		availablePermissions[permission] = true
	}
	return context.WithValue(ctx, ContextKeyPermissionHeader, availablePermissions)
}
