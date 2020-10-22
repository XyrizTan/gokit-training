package sample

import (
	"context"
	"net/http"

	endpoints "github.com/viki-org/gokit-template/pkg/endpoints/sample"
)

// DecodeRetrieveRequest decodes the http request into endpoints.{serviceName}.RetrieveRequest
func DecodeRetrieveRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req endpoints.RetrieveRequest

	// Get necessary parameters for req object from request route, body, headers, ...
	// vars := mux.Vars(r) // vars from request path
	// reqQuery := r.URL.Query()

	return req, nil
}
