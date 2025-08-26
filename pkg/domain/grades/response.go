package schools

type Response struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}
