package shared

type AddressDTO struct {
	ID             uint            `json:"id"`
	Street         string          `json:"street"`
	Number         string          `json:"number"`
	Complement     *string         `json:"complement"`
	ZipCode        string          `json:"zip_code"`
	NeighborhoodID uint            `json:"neighborhood_id"`
	Neighborhood   NeighborhoodDTO `json:"neighborhood"`
}
