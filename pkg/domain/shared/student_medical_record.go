package shared

import (
	"time"
)

type StudentMedicalRecordsDTO struct {
	ID            uint       `json:"id"`
	ConditionName string     `json:"condition_name"`
	DiagnosedAt   *time.Time `json:"diagnosed_at"`
	Notes         *string    `json:"notes"`
}
