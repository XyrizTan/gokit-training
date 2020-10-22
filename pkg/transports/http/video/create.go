package video

import (
	"context"
	"encoding/json"
	"net/http"

	endpoints "github.com/viki-org/gokit-template/pkg/endpoints/video"
)

// DecodeCreateRequest decodes the http request into endpoints.{serviceName}.RetrieveRequest
func DecodeCreateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req endpoints.CreateRequest

	// Get necessary parameters for req object from request route, body, headers, ...
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}
