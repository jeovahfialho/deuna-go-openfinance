package model

import "time"

// AuditLogEntry represents an entry in the audit trail
type AuditLogEntry struct {
	Timestamp   time.Time `json:"timestamp"`
	UserID      string    `json:"userId"`
	Action      string    `json:"action"`
	Description string    `json:"description"`
}
