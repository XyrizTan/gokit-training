package video

import (
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"

	videoModels "github.com/viki-org/gokit-template/pkg/models/video"
)

type dbStore struct {
	// We pass in the driver to the store here.
	// In this case dbStore we pass in gorm.DB
	// For example if redisStore we should pass in redis driver.
	db     *gorm.DB
	logger log.Logger
}

// NewDBStore returns a new db store instance.
func NewDBStore(db *gorm.DB, logger log.Logger) Store {
	return &dbStore{db, logger}
}

func (s *dbStore) Create(title string, duration int) (*videoModels.Video, error) {
	video := &videoModels.Video{
		Title:    title,
		Duration: duration,
	}
	errs := s.db.Create(video).GetErrors()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	return video, nil
}

func (s *dbStore) Retrieve(id int) (*videoModels.Video, error) {
	video := &videoModels.Video{}
	errs := s.db.Where("id = ?", id).First(video).GetErrors()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	return video, nil
}
