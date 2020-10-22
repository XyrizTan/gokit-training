package video

import (
	"github.com/go-kit/kit/log"

	videoModels "github.com/viki-org/gokit-template/pkg/models/video"
	videoStore "github.com/viki-org/gokit-template/pkg/stores/video"
)

// Service encapsulates business logic for login history.
type Service interface {
	Create(title string, duration int) (*videoModels.Video, error)
	Retrieve(id int) (*videoModels.Video, error)
}

type apiService struct {
	// Declare the dependencies for service inside.
	// Normally dependencies can be stores, logger, cache, queue producer, etc...
	store  videoStore.Store
	logger log.Logger
}

// NewAPIService returns apiService as a Service, which handles the business logic for the APIs.
func NewAPIService(store videoStore.Store, logger log.Logger) Service {
	return &apiService{store, logger}
}
