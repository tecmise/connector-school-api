package shared

type DirectorSchoolDTO struct {
	SchoolID  uint            `json:"school_id"`
	School    SchoolSimpleDTO `json:"school"`
	Status    string          `json:"status"`
	Year      int             `json:"year"`
	Notes     *string         `json:"notes"`
	StartDate string          `json:"start_date"`
	EndDate   *string         `json:"end_date"`
}
