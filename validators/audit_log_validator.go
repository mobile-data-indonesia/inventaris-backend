package validators

type CreateAuditLogRequest struct {
	AuditorID   string `json:"auditor_id" validate:"required"`
	AuditStatus string `json:"audit_status" validate:"required,oneof=good issue"`
	AuditNotes  string `json:"audit_notes" validate:"required"`
}