package school_user_profile

import (
	"github.com/tecmise/school-api-connector/pkg/domain/shared"
)

type SchoolTenantResponse struct {
	shared.TenantBaseResponse
	TenantId int64 `json:"school_id"`
}
