package main

type SearchByTitle struct {
	Query        string   `json:"query"`
	Page         int      `json:"page"`
	TotalRecords int      `json:"total_records"`
	TotalPages   int      `json:"total_pages"`
	Took         float64  `json:"took"`
	Records      []Record `json:"records"`
}

type Record struct {
	Date         int    `json:"date"`
	Filename     string `json:"filename"`
	Brief        Brief  `json:"brief"`
	JobNumber    string `json:"job_number"`
	UnitID       string `json:"unit_id"`
	UnitName     string `json:"unit_name"`
	UnitAPIURL   string `json:"unit_api_url"`
	TenderAPIURL string `json:"tender_api_url"`
}

type Brief struct {
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	Companies Companies `json:"companies"`
}

type Companies struct {
	IDs   []string `json:"ids"`
	Names []string `json:"names"`
}
