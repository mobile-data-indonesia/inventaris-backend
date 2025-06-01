package services

import (
	"fmt"

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

func (s *ItemService) GetItemByID(itemID uuid.UUID) (*models.Item, error) {
	var item models.Item
	if err := s.DB.Preload("Holder").First(&item, "item_id = ?", itemID).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *ItemService) UpdateItem(itemID uuid.UUID, req validators.UpdateItemRequest, imageURL string) error {
	var parsedHolderID *uuid.UUID
	if req.HolderID != nil {
		holderUUID, err := uuid.Parse(*req.HolderID)
		if err != nil {
			return fmt.Errorf("invalid holder_id: %w", err)
		}
		parsedHolderID = &holderUUID
	}

	item := &models.Item{
		ItemID:             itemID,
		HolderID:           parsedHolderID,
		ItemName:           req.ItemName,
		Vendor:             req.Vendor,
		Category:           req.Category,
		Location:           req.Location,
		ItemImageURL:       imageURL,
		ItemStatus:         req.ItemStatus,
		PurchaseDate:       req.PurchaseDate,
		InitialValue:       req.InitialValue,
		CurrentValue:       req.InitialValue, // Assuming current value is reset to initial value on update
		DepreciationRate:   req.DepreciationRate,
		DepreciationPeriod: req.DepreciationPeriod,
		ItemDescription:    req.ItemDescription,
	}

	return s.DB.Save(item).Error
}

func (s *ItemService) GetAllItems() ([]models.Item, error) {
	var items []models.Item
	if err := s.DB.Preload("Holder").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}