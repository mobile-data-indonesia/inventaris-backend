package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID          uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Username        string     `gorm:"size:50;unique;not null"`
	Password        string     `gorm:"size:100;not null"`
	Email           *string    `gorm:"size:100"`
	PhoneNumber     *string    `gorm:"size:20"`
	Title           string     `gorm:"size:50;not null"`
	Role            string     `gorm:"size:20;not null"`
	Department      string     `gorm:"size:50;not null"`
	ProfileImageURL string     `gorm:"size:255"`
	DeletedAt       *time.Time `gorm:"index"`
	CreatedAt       time.Time  `gorm:"autoCreateTime"`
	UpdatedAt       time.Time  `gorm:"autoUpdateTime"`

	AuditLogs      []AuditLog  `gorm:"foreignKey:AuditorID;references:UserID"`
	Items          []Item  `gorm:"foreignKey:HolderID;references:UserID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UserID = uuid.New()
	u.ProfileImageURL = "default-profile.png"

	return nil
}
