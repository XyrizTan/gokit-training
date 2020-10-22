package video

// Video is the mapping of video model with video record in the DB.
type Video struct {
	ID       int    `gorm:"column:id;primary_key;auto_increment"`
	Title    string `gorm:"column:title"`
	Duration int    `gorm:"column:duration"`
}

func (u Video) TableName() string {
	return "videos"
}
