package shared

type CountryDTO struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	IsoCode   string `json:"iso_code"`
	PhoneCode string `json:"phone_code"`
}
