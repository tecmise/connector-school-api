package cluster_user_profile

type Response struct {
	ClusterId int64 `json:"cluster_id"`
	UserId    int64 `json:"user_id"`
	ProfileId int64 `json:"profile_id"`
}
