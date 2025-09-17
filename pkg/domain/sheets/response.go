package sheets

type Response struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	URL           string `json:"url"`
	Status        string `json:"status"`
	Total         int    `json:"total"`
	SuccessTotal  int    `json:"success_total"`
	ErrorTotal    int    `json:"error_total"`
	ImportType    string `json:"import_type"`
	Origin        string `json:"origin"`
	Parameters    string `json:"parameters"`
	SchoolID      uint   `json:"school_id"`
	ClassID       uint   `json:"class_id"`
	IsIncremental bool   `json:"is_incremental"`
}
