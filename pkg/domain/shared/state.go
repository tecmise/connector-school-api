package shared

type StateDTO struct {
	ID           uint       `json:"id"`
	Name         string     `json:"name"`
	Abbreviation string     `json:"abbreviation"`
	CountryID    uint       `json:"country_id"`
	Country      CountryDTO `json:"country"`
}
