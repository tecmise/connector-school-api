package shared

type StudentSchoolClassDTO struct {
	StudentSchoolID uint             `json:"student_school_id"`
	StudentSchool   StudentSchoolDTO `json:"-"`
	ClassID         uint             `json:"class_id"`
	Class           ClassSimpleDTO   `json:"class"`
	Status          string           `json:"status"` // 'enrolled', 'graduated', 'finished'
	Year            int              `json:"year"`
	StartDate       string           `json:"start_date"`
	EndDate         *string          `json:"end_date"`
}
