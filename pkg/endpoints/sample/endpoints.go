package sample

import (
	"github.com/go-kit/kit/endpoint"
	sampleService "github.com/viki-org/gokit-template/pkg/services/sample"
)

// Endpoints encapsulate login history endpoints.
// This layer is for input/output validation.
type Endpoints struct {
	RetrieveEndpoint endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the provided service.
func MakeServerEndpoints(s sampleService.Service) Endpoints {
	return Endpoints{
		RetrieveEndpoint: RetrieveEndpoint(s),
	}
}
