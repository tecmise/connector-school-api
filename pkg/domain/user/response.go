package schools

type Response struct {
	ID       string  `json:"id"`
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	IsActive bool    `json:"is_active"`
}
