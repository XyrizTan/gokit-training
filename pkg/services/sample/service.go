package sample

import (
	"github.com/go-kit/kit/log"
	sampleStore "github.com/viki-org/gokit-template/pkg/stores/sample"
)

// Service encapsulates business logic for login history.
type Service interface {
	Retrieve() (int, error)
}

type apiService struct {
	// Declare the dependencies for service inside.
	// Normally dependencies can be stores, logger, cache, queue producer, etc...
	store  sampleStore.Store
	logger log.Logger
}

// NewAPIService returns apiService as a Service, which handles the business logic for the APIs.
func NewAPIService(store sampleStore.Store, logger log.Logger) Service {
	return &apiService{store, logger}
}
