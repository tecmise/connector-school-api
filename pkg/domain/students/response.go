package students

import (
	"github.com/tecmise/connector-school-api/pkg/domain/shared"
)

type Response struct {
	ID                    uint                           `json:"id"`
	PersonID              uint                           `json:"person_id"`
	Person                shared.PersonDTO               `json:"person"`
	Registration          string                         `json:"registration"`
	NameOfResponsible     string                         `json:"name_of_responsible"`
	DocumentOfResponsible string                         `json:"document_of_responsible"`
	EmailOfResponsible    string                         `json:"email_of_responsible"`
	PhoneOfResponsible    string                         `json:"phone_of_responsible"`
	StudentSchools        []shared.StudentSchoolDTO      `json:"student_schools"`
	StudentOccurrences    []shared.StudentOccurrencesDTO `json:"student_occurrences"`
	IsActive              bool                           `json:"is_active"`
}
