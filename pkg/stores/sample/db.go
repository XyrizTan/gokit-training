package sample

import (
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
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

func (s *dbStore) Retrieve() (int, error) {
	// Put the db store operation here
	return 1, nil
}
