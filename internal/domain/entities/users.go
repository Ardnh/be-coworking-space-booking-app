package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	UserID       uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Email        string         `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password     string         `gorm:"type:varchar(255);not null"`
	Username     string         `gorm:"type:varchar(255);index;not null"`
	FullName     string         `gorm:"type:varchar(255);not null"`
	PhoneNumber  string         `gorm:"type:varchar(20);not null"`
	UserType     string         `gorm:"type:varchar(20);not null"`
	ProfileImage string         `gorm:"type:varchar(255);not null"`
	Status       string         `gorm:"type:varchar(20);not null;default:'active'"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

// TableName overrides the table name used by Users to `users`
func (Users) TableName() string {
	return "users"
}
