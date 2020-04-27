package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Post : Post Model
type Post struct {
	gorm.Model
	ID uint64 `gorm:"primary_key"` // auto increment uint default true

	Title    string
	Content  string
	Username string
	Date     time.Time
}
