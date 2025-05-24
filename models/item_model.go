package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	ItemID             uuid.UUID      `gorm:"type:uuid;primaryKey"`
	HolderID           *uuid.UUID     `gorm:"type:uuid"`
	Holder             *User          `gorm:"foreignKey:HolderID"`
	ItemName           string         `gorm:"type:varchar(100);not null"`
	Vendor             string         `gorm:"type:varchar(100);not null"`
	Category           string         `gorm:"type:varchar(100);not null"`
	Location           string         `gorm:"type:varchar(100);not null"`
	ItemImageURL       string         `gorm:"type:varchar(255);not null"`
	ItemStatus         string         `gorm:"type:varchar(20);not null"`
	PurchaseDate       time.Time      `gorm:"type:timestamp;notnull"`
	InitialValue       int 				`gorm:"type:int;notnull"`
	CurrentValue       int 				`gorm:"type:int;notnull"`
	DepreciationRate   float64 			`gorm:"type:float;notnull"`
	DepreciationPeriod string         `gorm:"type:varchar(20) ;not null"`
	ItemDescription    string         `gorm:"type:text"`
	DeletedAt          *time.Time 		`gorm:"index"` 
	CreatedAt          time.Time      `gorm:"autoCreateTime"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime"`
}

func (u *Item) BeforeCreate(tx *gorm.DB) (err error) {
	u.ItemID = uuid.New()
	u.ItemImageURL = "default-profile.png"

	return nil
}
