package restricted_persons

import (
	"github.com/tecmise/connector-school-api/pkg/domain/schools"
	"github.com/tecmise/connector-school-api/pkg/domain/shared"
)

type Response struct {
	ID          uint               `json:"id"`
	FullName    string             `json:"full_name"`
	Gender      string             `json:"gender"`
	ImageID     uint               `json:"image_id"`
	Image       shared.ImageDTO    `json:"image"`
	DocumentID  *uint              `json:"document_id"`
	Document    shared.DocumentDTO `json:"document"`
	Description string             `json:"description"`
	IsActive    bool               `json:"is_active"`
	Schools     []schools.Response `json:"schools"`
}
