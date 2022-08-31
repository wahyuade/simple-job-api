package models

type Job struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	Url         string `json:"url"`
	CreatedAt   string `json:"created_at"`
	CompanyUrl  string `json:"company_url"`
	Location    string `json:"location"`
	Title       string `json:"title"`
	Description string `json:"description"`
	HowToApply  string `json:"how_to_apply"`
	CompanyLogo string `json:"company_logo"`
}
