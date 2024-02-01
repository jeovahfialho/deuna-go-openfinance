package model

import "time"

type AuditLog struct {
	ID        string
	Action    string
	UserID    string
	Timestamp time.Time
	Details   string
}
