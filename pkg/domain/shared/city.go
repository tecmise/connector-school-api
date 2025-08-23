package shared

type CityDTO struct {
	ID      uint     `json:"id"`
	Name    string   `json:"name"`
	StateID uint     `json:"state_id"`
	State   StateDTO `json:"state"`
}
