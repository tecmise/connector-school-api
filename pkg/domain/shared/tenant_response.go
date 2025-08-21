package shared

import "time"

type TenantBaseResponse struct {
	UserId     int64     `json:"user_id"`
	ProfileId  int64     `json:"profile_id"`
	ExpiressIn time.Time `json:"expires_in"`
}
