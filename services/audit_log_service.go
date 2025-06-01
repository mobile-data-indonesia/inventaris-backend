package services

import (
	"github.com/google/uuid"
	"github.com/mobile-data-indonesia/inventaris-backend/models"
	"github.com/mobile-data-indonesia/inventaris-backend/validators"
	"gorm.io/gorm"
)

type AuditLogService struct {
    DB *gorm.DB
}

func NewAuditLogService(db *gorm.DB) *AuditLogService {
    return &AuditLogService{DB: db}
}

func (s *AuditLogService) GetAllAuditLogs() ([]models.AuditLog, error) {
    var logs []models.AuditLog
    if err := s.DB.Preload("Auditor").Find(&logs).Error; err != nil {
        return nil, err
    }
    return logs, nil
}

func (s *AuditLogService) CreateAuditLog(input validators.CreateAuditLogRequest) error {
    auditorUUID, err := uuid.Parse(input.AuditorID)
    if err != nil {
        return err
    }
    auditLog := models.AuditLog{
        AuditID:      uuid.New(),
        AuditorID:    auditorUUID,
        AuditStatus:  input.AuditStatus,
        AuditNotes:   input.AuditNotes,
    }
    return s.DB.Create(&auditLog).Error
}