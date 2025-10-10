package students

import (
	"github.com/tecmise/connector-school-api/pkg/domain/shared"
)

type Response struct {
	ID                 uint                           `json:"id"`
	PersonID           uint                           `json:"person_id"`
	Person             shared.PersonDTO               `json:"person"`
	Responsibles       []shared.StudentResponsibleDTO `json:"responsibles"`
	StudentSchools     []shared.StudentSchoolDTO      `json:"student_schools"`
	StudentOccurrences []shared.StudentOccurrencesDTO `json:"student_occurrences"`
	IsActive           bool                           `json:"is_active"`
}
