package shared

type StudentFamilyMembersDTO struct {
	ID       uint    `json:"id"`
	Relation string  `json:"relation"`
	Name     string  `json:"name"`
	Phone    *string `json:"phone"`
	Notes    *string `json:"notes"`
}
