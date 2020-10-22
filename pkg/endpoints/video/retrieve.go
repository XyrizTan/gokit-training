package video

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	videoService "github.com/viki-org/gokit-template/pkg/services/video"
)

// RetrieveRequest encapsulates request payload structure
type RetrieveRequest struct {
	ID int
}

// RetrieveResponse encapsulates response payload structure
type RetrieveResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Duration int    `json:"duration"`
}

// RetrieveEndpoint handles input/output validation for a retrieve endpoint.
func RetrieveEndpoint(s videoService.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(RetrieveRequest)
		// Do the input valiation with req here.
		// E.g. validateVideoID(req.ID)

		video, err := s.Retrieve(req.ID)
		if err != nil {
			return nil, err
		}

		resp := RetrieveResponse{
			ID:       video.ID,
			Title:    video.Title,
			Duration: video.Duration,
		}

		return resp, nil
	}
}
