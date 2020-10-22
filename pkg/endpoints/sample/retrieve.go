package sample

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	sampleService "github.com/viki-org/gokit-template/pkg/services/sample"
)

// RetrieveRequest encapsulates request payload structure
type RetrieveRequest struct {
}

// RetrieveResponse encapsulates response payload structure
type RetrieveResponse struct {
	Value int `json:"value"`
}

// RetrieveEndpoint handles input/output validation for a retrieve endpoint.
func RetrieveEndpoint(s sampleService.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// req := request.(RetrieveRequest)
		// Do the input valiation with req here.

		val, err := s.Retrieve()
		if err != nil {
			return nil, err
		}

		resp := RetrieveResponse{Value: val}

		return resp, nil
	}
}
