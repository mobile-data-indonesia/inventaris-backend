package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/mobile-data-indonesia/inventaris-backend/services"
	"github.com/mobile-data-indonesia/inventaris-backend/validators"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ItemHandler struct {
	ItemService *services.ItemService
}

func NewItemHandler(s *services.ItemService) *ItemHandler {
	return &ItemHandler{ItemService: s}
}

func (h *ItemHandler) CreateItem(c *gin.Context) {
	log.Println("CreateItem called")

	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse multipart form"})
		return
	}

	var input validators.CreateItemRequest
	if err := c.ShouldBindWith(&input, binding.FormMultipart); err != nil {
		log.Println("Form bind error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Buat UUID untuk item
	itemID := uuid.New()

	var itemImageURL string

	// Ambil file item_image
	file, err := c.FormFile("item_image")
	if err != nil {
		itemImageURL = "uploads/items/placeholder.png"
	} else {
	
		ext := filepath.Ext(file.Filename)
		fileName := itemID.String() + ext
		dstPath := filepath.Join("uploads/items", fileName)

		// Pastikan folder uploads/items ada
		if err := os.MkdirAll("uploads/items", os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create upload directory"})
			return
		}

		// Simpan file
		if err := c.SaveUploadedFile(file, dstPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save image"})
			return
		}

		// Buat URL atau path relatif
		itemImageURL = "uploads/items/" + fileName
	}

	// Call service untuk simpan ke DB
	if err := h.ItemService.CreateItem(input, itemID, itemImageURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "item created successfully"})
}

func (h *ItemHandler) GetItemByID(c *gin.Context) {
	log.Println("GetItemByID called")

	itemIDStr := c.Param("id")
	itemID, err := uuid.Parse(itemIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}

	item, err := h.ItemService.GetItemByID(itemID)
	if
	 err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) UpdateItem(c *gin.Context) {
	log.Println("UpdateItem called")

	itemIDStr := c.Param("id")
	itemID, err := uuid.Parse(itemIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}

	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse multipart form"})
		return
	}

	var input validators.UpdateItemRequest
	if err := c.ShouldBindWith(&input, binding.FormMultipart); err != nil {
		log.Println("Form bind error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var itemImageURL string

	file, err := c.FormFile("item_image")
	if err == nil {
		ext := filepath.Ext(file.Filename)
		fileName := itemID.String() + ext
		dstPath := filepath.Join("uploads/items", fileName)

		if err := os.MkdirAll("uploads/items", os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create upload directory"})
			return
		}

		if err := c.SaveUploadedFile(file, dstPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save image"})
			return
		}

		itemImageURL = "uploads/items/" + fileName
	} else {
		itemImageURL = "uploads/items/placeholder.png"
	}

	if err := h.ItemService.UpdateItem(itemID, input, itemImageURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "item updated successfully"})
}

func (h *ItemHandler) GetAllItems(c *gin.Context) {
	log.Println("GetAllItems called")

	items, err := h.ItemService.GetAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}
