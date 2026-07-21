package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"Name" gorm:"type:varchar(255);not null"`
	Email        string `json:"Email" gorm:"type:varchar(255);not null"`
	PasswordHash string `json:"-" gorm:"not null"`
	RoleID       int    `json:"role_id" gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
