package services

import (
	"github.com/google/uuid"
	"github.com/mobile-data-indonesia/inventaris-backend/models"
	"github.com/mobile-data-indonesia/inventaris-backend/validators"
	"gorm.io/gorm"
)

type ItemService struct {
	DB *gorm.DB
}

func NewItemService(db *gorm.DB) *ItemService {
	return &ItemService{DB: db}
}

func (s *ItemService) CreateItem(req validators.CreateItemRequest, itemID uuid.UUID, imageURL string) error {
	item := &models.Item{
		ItemID:             itemID,
		HolderID:           req.HolderID,
		ItemName:           req.ItemName,
		Vendor:             req.Vendor,
		Category:           req.Category,
		Location:           req.Location,
		ItemImageURL:       imageURL,
		ItemStatus:         req.ItemStatus,
		PurchaseDate:       req.PurchaseDate,
		InitialValue:       req.InitialValue,
		CurrentValue:       req.InitialValue,
		DepreciationRate:   req.DepreciationRate,
		DepreciationPeriod: req.DepreciationPeriod,
		ItemDescription:    req.ItemDescription,
	}

	return s.DB.Create(item).Error
}
