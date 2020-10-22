package video

import (
	videoModels "github.com/viki-org/gokit-template/pkg/models/video"
)

// Store encapsulates operations to a data store (e.g. postgres, redis, ...).
type Store interface {
	Create(title string, duration int) (*videoModels.Video, error)
	Retrieve(id int) (*videoModels.Video, error)
}
