package models

import (
	"time"

	"gorm.io/gorm"
)

type Url struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Path        string `json:"path" gorm:"type:varchar(255);not null;uniqueIndex"`
	Destination string `json:"destination" gorm:"type:varchar(2048);not null"`
	UserId      uint   `json:"user_id" gorm:"not null"`
	User        User   `json:"-" gorm:"foreignKey:UserId;references:ID;constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
