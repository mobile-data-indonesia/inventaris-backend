package validators

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type CreateItemRequest struct {
	HolderID           *uuid.UUID              `form:"holder_id"` // optional
	ItemName           string                  `form:"item_name" validate:"required"`
	Vendor             string                  `form:"vendor" validate:"required"`
	Category           string                  `form:"category" validate:"required"`
	Location           string                  `form:"location" validate:"required"`
	ItemImage          *multipart.FileHeader   `form:"item_image" validate:"required"` // image file upload
	ItemStatus         string                  `form:"item_status" validate:"required,oneof=good lost damaged"`
	PurchaseDate       time.Time               `form:"purchase_date" validate:"required"`
	InitialValue       int                     `form:"initial_value" validate:"required,min=0"`
	CurrentValue       int                     `form:"current_value" validate:"required,min=0"`
	DepreciationRate   float64                 `form:"depreciation_rate" validate:"required,gte=0"`
	DepreciationPeriod string                  `form:"depreciation_period" validate:"required,oneof=monthly yearly"`
	ItemDescription    string                  `form:"item_description"` // optional
}