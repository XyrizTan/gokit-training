package video

import (
	videoModels "github.com/viki-org/gokit-template/pkg/models/video"
)

func (s *apiService) Retrieve(id int) (*videoModels.Video, error) {
	// Do your business logic in this method.
	//
	return s.store.Retrieve(id)
}
