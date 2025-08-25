package shared

type ImageDTO struct {
	ID       uint    `json:"id"`
	Path     string  `json:"path"`
	Size     *string `json:"size"`
	Format   *string `json:"format"`
	IsActive bool    `json:"is_active"`
}
