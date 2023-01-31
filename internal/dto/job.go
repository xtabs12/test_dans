package dto

type (
	GetJobListParams struct {
		Page               int    `json:"page"`
		SearchDescription  string `json:"search_description"`
		SearchLocation     string `json:"search_location"`
		SearchOnlyFullTime bool   `json:"search_only_full_time"`
	}
	GetJobDetailParams struct {
		ID string `json:"id"`
	}
	Job struct {
		ID          string `json:"id"`
		Type        string `json:"type"`
		URI         string `json:"url"`
		CreatedAt   string `json:"created_at"`
		Company     string `json:"company"`
		CompanyURI  string `json:"company_url"`
		Location    string `json:"location"`
		Title       string `json:"title"`
		Description string `json:"description"`
		HowToApply  string `json:"how_to_apply"`
		CompanyLogo string `json:"company_logo"`
	}
)

func NewJob() *Job {
	return new(Job)
}
