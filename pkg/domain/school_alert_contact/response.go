package school_alert_contact

import (
	"github.com/tecmise/connector-school-api/pkg/domain/shared"
)

type Response struct {
	Person       shared.PersonDTO         `json:"person"`
	SchoolsRoles []shared.SchoolRoleGroup `json:"schools_roles"`
}
