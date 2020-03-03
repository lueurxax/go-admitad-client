package responses

type News struct {
	Results []ResultsNews `json:"results"`
	Meta    Meta          `json:"_meta"`
}
type Advcampaign struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type ResultsNews struct {
	Language     string      `json:"language"`
	Content      string      `json:"content"`
	ID           int         `json:"id"`
	URL          string      `json:"url"` // Ссылка на сторонний ресурс
	ShortContent string      `json:"short_content"`
	Advcampaign  Advcampaign `json:"advcampaign,omitempty"`
	Datetime     string      `json:"datetime"`
}
