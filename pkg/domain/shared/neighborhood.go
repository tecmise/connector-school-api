package shared

type NeighborhoodDTO struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	CityID uint    `json:"city_id"`
	City   CityDTO `json:"city"`
}
