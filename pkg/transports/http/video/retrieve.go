package video

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	endpoints "github.com/viki-org/gokit-template/pkg/endpoints/video"
	"github.com/viki-org/gokit-template/pkg/transports/http/errors"
)

// DecodeRetrieveRequest decodes the http request into endpoints.{serviceName}.RetrieveRequest
func DecodeRetrieveRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req endpoints.RetrieveRequest

	// Get necessary parameters for req object from request route, body, headers, ...
	vars := mux.Vars(r) // vars from request path
	videoID, ok := vars["videoID"]
	if !ok {
		return nil, errors.ErrInvalidRequest
	}
	videoIDint, err := strconv.Atoi(videoID)
	if err != nil {
		return nil, errors.ErrInvalidRequest
	}
	req.ID = videoIDint

	return req, nil
}
