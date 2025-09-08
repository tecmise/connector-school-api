package shared

import (
	"time"
)

type PersonDTO struct {
	ID          uint       `json:"id"`
	ImageID     *uint      `json:"image_id"`
	Image       *ImageDTO  `json:"image"`
	ShieldID    *string    `json:"shield_id"`
	Email       *string    `json:"email"`
	Document    string     `json:"document"`
	FullName    string     `json:"full_name"`
	BirthDate   time.Time  `json:"birth_date"`
	Gender      string     `json:"gender"`
	PhoneNumber *string    `json:"phone_number"`
	AddressID   uint       `json:"address_id"`
	Address     AddressDTO `json:"address"`
	IsActive    bool       `json:"is_active"`
}
