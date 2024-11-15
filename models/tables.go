package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;type:char(36)"`
	Firstname string
	Lastname  string
	Username  string `gorm:"unique"`
	Password  string
	Type      string `gorm:"default:user"`
	CreatedAt time.Time
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate a new UUID and assign it to the ID field
	user.ID = uuid.New()
	return nil
}
