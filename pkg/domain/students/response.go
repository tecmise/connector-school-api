package students

import "github.com/tecmise/connector-school-api/pkg/domain/shared"

type Response struct {
	ID                    uint             `json:"id"`
	PersonID              uint             `json:"person_id"`
	Person                shared.PersonDTO `json:"person"`
	Registration          string           `json:"registration"`
	NameOfResponsible     string           `json:"name_of_responsible"`
	DocumentOfResponsible string           `json:"document_of_responsible"`
	EmailOfResponsible    string           `json:"email_of_responsible"`
	ClassID               uint             `json:"class_id"`
	// Class                 ClassDTO                          `json:"class"`
	StudentFamilyMembers  []shared.StudentFamilyMembersDTO  `json:"student_family_members"`
	StudentMedicalRecords []shared.StudentMedicalRecordsDTO `json:"student_medical_records"`
	StudentOccurrences    []shared.StudentOccurrencesDTO    `json:"student_occurrences"`
	IsActive              bool                              `json:"is_active"`
}
