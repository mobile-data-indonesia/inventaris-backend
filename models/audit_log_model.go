package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuditLog struct {
	AuditID         uuid.UUID  `gorm:"type:uuid;primaryKey"`
	AuditorID       uuid.UUID  `gorm:"type:uuid;not null"`
	Auditor         User      	`gorm:"foreignKey:AuditorID;references:UserID"`
	AuditStatus   	string     `gorm:"type:varchar(255);not null"` 
	AuditNotes      string     `gorm:"type:text;not null"`
	DeletedAt       *time.Time `gorm:"index"`
	CreatedAt       time.Time  `gorm:"autoCreateTime"`
	UpdatedAt       time.Time  `gorm:"autoUpdateTime"`
}


func (u *AuditLog) BeforeCreate(tx *gorm.DB) (err error) {
	u.AuditID = uuid.New()

	return nil
}
