package shared

import (
	"github.com/tecmise/connector-school-api/pkg/domain/subjects"
)

type ClassSimpleDTO struct {
	ID       uint                `json:"id"`
	Name     string              `json:"name"`
	GradeID  uint                `json:"grade_id"` // ano escolar (1, 2, 3, etc)
	Grade    GradeDTO            `json:"grade"`
	ShiftID  uint                `json:"shift_id"` // turno escolar (manha, tarde ou noite)
	Shift    ShiftDTO            `json:"shift"`
	LevelID  uint                `json:"level_id"` // infantil, medio ou fundamental
	Level    LevelDTO            `json:"level"`
	SchoolID uint                `json:"school_id"`
	School   SchoolSimpleDTO     `json:"school"`
	IsActive bool                `json:"is_active"`
	Subjects []subjects.Response `json:"subjects"`
}
