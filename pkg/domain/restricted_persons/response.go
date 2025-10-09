package restricted_persons

import (
	"github.com/tecmise/connector-school-api/pkg/domain/schools"
	"github.com/tecmise/connector-school-api/pkg/domain/shared"
)

type Response struct {
	ID          uint               `json:"id"`
	Person      shared.PersonDTO   `json:"person"`
	Type        string             `json:"type"`
	Description string             `json:"description"`
	IsActive    bool               `json:"is_active"`
	Schools     []schools.Response `json:"schools"`
	Images      []shared.ImageDTO  `json:"images"`
}
