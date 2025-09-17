package shared

type StudentResponsibleDTO struct {
	ResponsibleID uint           `json:"responsible_id"`
	Responsible   ResponsibleDTO `json:"responsible"`
}
