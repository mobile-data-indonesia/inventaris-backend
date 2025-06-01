package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mobile-data-indonesia/inventaris-backend/services"
	"github.com/mobile-data-indonesia/inventaris-backend/validators"
)

type AuditLogHandler struct {
    AuditLogService *services.AuditLogService
}

func NewAuditLogHandler(s *services.AuditLogService) *AuditLogHandler {
    return &AuditLogHandler{AuditLogService: s}
}

func (h *AuditLogHandler) GetAllAuditLogs(c *gin.Context) {
    logs, err := h.AuditLogService.GetAllAuditLogs()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, logs)
}

func (h *AuditLogHandler) CreateAuditLog(c *gin.Context) {
	var input validators.CreateAuditLogRequest
	log.Printf("Input: %+v", input)
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.AuditLogService.CreateAuditLog(input); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "audit log created successfully"})
}