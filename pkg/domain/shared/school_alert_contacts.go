package shared

type SchoolRoleGroup struct {
	Role    string        `json:"role"`
	Schools []SchoolBasic `json:"schools"`
}

type SchoolBasic struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
