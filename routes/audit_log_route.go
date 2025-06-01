package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mobile-data-indonesia/inventaris-backend/handlers"
)

type AuditLogRoutes struct {
	AuditLogController *handlers.AuditLogHandler
}

func NewAuditLogRoutes(ctrl *handlers.AuditLogHandler) *AuditLogRoutes {
	return &AuditLogRoutes{AuditLogController: ctrl}
}

func (r *AuditLogRoutes) RegisterRoutes(router *gin.Engine) {
	audit := router.Group("/audit-logs")
	{
		audit.GET("/", r.AuditLogController.GetAllAuditLogs)
		audit.POST("/", r.AuditLogController.CreateAuditLog)
	}
}

