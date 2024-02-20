package utilities

type BookInfo struct {
	BookCount int 		 `json:"count"`
	Authors []string    `json:"authors,omitempty"`
	Langugage []string  `json:"language"`
}

type TotalBookCount struct {
	Count int `json:"count"`
}

type BookCounter struct {
	Langugage string  `json:"language"`
	Books int 		  `json:"books"`
	Authors int 	  `json:"authors,omitempty"`
	Fraction float64  `json:"fraction"`
}

type Status struct {
	Gutendexapi int  `json:"gutendexapi"`
	Langugaeapi int  `json:"languageapi"`
	Countriesapi int `json:"countriesapi"`
	Version string   `json:"version"`
	Uptime float64   `json:"uptime"`
}
