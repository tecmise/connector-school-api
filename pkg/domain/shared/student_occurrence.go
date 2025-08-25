package shared

import (
	"time"
)

type StudentOccurrencesDTO struct {
	ID             uint      `json:"id"`
	OccurrenceDate time.Time `json:"occurrence_date"`
	Description    string    `json:"description"`
	Severity       *string   `json:"severity"`
	ActionTaken    *string   `json:"action_taken"`
}
