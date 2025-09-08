package shared

type SchoolSimpleDTO struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Document  string  `json:"document"`
	Phone     *string `json:"phone"`
	Email     *string `json:"email"`
	AddressID uint    `json:"address_id"`
	IsActive  bool    `json:"is_active"`
}
