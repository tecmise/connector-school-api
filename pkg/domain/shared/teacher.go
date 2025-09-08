package shared

import (
	"github.com/tecmise/connector-school-api/pkg/domain/users"
)

type TeacherSimpleDTO struct {
	ID       uint            `json:"id"`
	UserID   *string         `json:"user_id"`
	User     *users.Response `json:"user"`
	PersonID uint            `json:"person_id"`
	Person   PersonSimpleDTO `json:"person"`
	IsActive bool            `json:"is_active"`
}
