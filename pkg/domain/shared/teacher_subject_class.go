package shared

import (
	"github.com/tecmise/connector-school-api/pkg/domain/subjects"
)

type TeacherSubjectClassDTO struct {
	TeacherID uint              `json:"teacher_id"`
	Teacher   TeacherSimpleDTO  `json:"-"`
	SubjectID uint              `json:"subject_id"`
	Subject   subjects.Response `json:"subject"`
	ClassID   uint              `json:"class_id"`
	Class     ClassSimpleDTO    `json:"class"`
	Year      string            `json:"year"`
	StartDate string            `json:"start_date"`
	EndDate   *string           `json:"end_date"`
}
