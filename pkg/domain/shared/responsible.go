package shared

type ResponsibleDTO struct {
	ID       uint      `json:"id"`
	Relation string    `json:"status"` // 'father', 'mother', 'brother', 'other'
	PersonID uint      `json:"person_id"`
	Person   PersonDTO `json:"person"`
	// Students []StudentDTO `json:"students"`
}
