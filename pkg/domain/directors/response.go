package directors

import (
	"github.com/tecmise/connector-school-api/pkg/domain/shared"
	"github.com/tecmise/connector-school-api/pkg/domain/users"
)

type Response struct {
	ID       uint             `json:"id"`
	UserID   *string          `json:"user_id"`
	User     users.Response   `json:"user"`
	PersonID uint             `json:"person_id"`
	Person   shared.PersonDTO `json:"person"`
	IsActive bool             `json:"is_active"`

	Schools []shared.DirectorSchoolDTO `json:"schools"`
}
