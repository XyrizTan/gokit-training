package video

import (
	videoModels "github.com/viki-org/gokit-template/pkg/models/video"
)

func (s *apiService) Create(title string, duration int) (*videoModels.Video, error) {
	// Do your business logic in this method.
	// adjustSmartVideoDuration(duration)
	video, err := s.store.Create(title, duration)
	if err != nil {
		return nil, err
	}
	// updateVideoResolution(video.ID)
	return video, nil
}
