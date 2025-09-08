package restricted_persons

import (
	"github.com/tecmise/connector-school-api/pkg/domain/schools"
	"github.com/tecmise/connector-school-api/pkg/domain/shared"
)

type Response struct {
	ID          int                `json:"id"`
	PersonID    int                `json:"person_id"`
	Person      shared.PersonDTO   `json:"person"`
	Description string             `json:"description"`
	IsActive    bool               `json:"is_active"`
	Schools     []schools.Response `json:"schools"`
}
