package video

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	videoService "github.com/viki-org/gokit-template/pkg/services/video"
)

// CreateRequest encapsulates request payload structure
type CreateRequest struct {
	Title    string `json:"title"`
	Duration int    `json:"duration"`
}

// CreateResponse encapsulates response payload structure
type CreateResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Duration int    `json:"duration"`
}

// CreateEndpoint handles input/output validation for a retrieve endpoint.
func CreateEndpoint(s videoService.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateRequest)
		// Do the input valiation with req here.
		// E.g. validateVideoTitle(req.Title), validateVideoDuration(req.Duration)

		video, err := s.Create(req.Title, req.Duration)
		if err != nil {
			return nil, err
		}

		resp := CreateResponse{
			ID:       video.ID,
			Title:    video.Title,
			Duration: video.Duration,
		}

		return resp, nil
	}
}
