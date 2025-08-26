package classes

import (
	"github.com/tecmise/connector-school-api/pkg/domain/schools"
	"github.com/tecmise/connector-school-api/pkg/domain/shared"
)

type Response struct {
	ID       uint             `json:"id"`
	Name     string           `json:"name"`
	GradeID  uint             `json:"grade_id"` // ano escolar (1, 2, 3, etc)
	Grade    shared.GradeDTO  `json:"grade"`
	ShiftID  uint             `json:"shift_id"` // turno escolar (manha, tarde ou noite)
	Shift    shared.ShiftDTO  `json:"shift"`
	LevelID  uint             `json:"level_id"` // infantil, medio ou fundamental
	Level    shared.LevelDTO  `json:"level"`
	SchoolID uint             `json:"school_id"`
	School   schools.Response `json:"school"`
	IsActive bool             `json:"is_active"`
}
