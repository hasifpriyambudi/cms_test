package domain

import (
	"time"

	"gorm.io/gorm"
)

type News struct {
	Id        uint   `json:"id"`
	Name      string `gorm:"primary_key" json:"name"`
	Username  string `gorm:"size:30;not null;uniqueIndex" json:"username"`
	Password  string `gorm:"size:255;not null;uniqueIndex" json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
