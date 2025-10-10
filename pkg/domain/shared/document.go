package shared

type DocumentDTO struct {
	ID       uint    `json:"id"`
	Path     string  `json:"path"`
	Size     *string `json:"size"`
	Format   *string `json:"format"`
	IsActive bool    `json:"is_active"`
}
