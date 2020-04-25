package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Post : Post Model
type Post struct {
	gorm.Model
	Title    string
	Username string
	Date     time.Time
}
