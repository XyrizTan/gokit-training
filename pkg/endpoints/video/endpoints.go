package video

import (
	"github.com/go-kit/kit/endpoint"

	videoService "github.com/viki-org/gokit-template/pkg/services/video"
)

// Endpoints encapsulate login history endpoints.
// This layer is for input/output validation.
type Endpoints struct {
	CreateEndpoint   endpoint.Endpoint
	RetrieveEndpoint endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the provided service.
func MakeServerEndpoints(s videoService.Service) Endpoints {
	return Endpoints{
		CreateEndpoint:   CreateEndpoint(s),
		RetrieveEndpoint: RetrieveEndpoint(s),
	}
}
