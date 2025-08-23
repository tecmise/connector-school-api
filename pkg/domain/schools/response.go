package schools

import "github.com/tecmise/connector-school-api/pkg/domain/shared"

type Response struct {
	ID        uint              `json:"id"`
	Name      string            `json:"name"`
	Type      string            `json:"type"`
	AddressID uint              `json:"address_id"`
	Address   shared.AddressDTO `json:"address"`
	IsActive  bool              `json:"is_active"`
}
