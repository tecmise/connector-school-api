package shared

import "github.com/tecmise/connector-school-api/pkg/domain/subjects"

type ClassSimpleDTO struct {
	ID       uint                `json:"id"`
	Name     string              `json:"name"`
	GradeID  uint                `json:"grade_id"`
	ShiftID  uint                `json:"shift_id"`
	LevelID  uint                `json:"level_id"`
	SchoolID uint                `json:"school_id"`
	IsActive bool                `json:"is_active"`
	Subjects []subjects.Response `json:"subjects"`
}
