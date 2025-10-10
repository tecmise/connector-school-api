package shared

type StudentSchoolDTO struct {
	ID                 uint            `json:"id"`
	StudentID          uint            `json:"student_id"`
	SchoolID           uint            `json:"school_id"`
	School             SchoolSimpleDTO `json:"school"`
	RegistrationNumber string          `json:"registration_number"`
	Status             string          `json:"status"` // 'enrolled', 'graduated', 'finished'
	Year               int             `json:"year"`
	StartDate          string          `json:"start_date"`
	EndDate            *string         `json:"end_date"`

	StudentSchoolClasses []StudentSchoolClassDTO `json:"student_school_classes"`
}
