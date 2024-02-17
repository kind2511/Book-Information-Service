package utilities

type BookInfo struct {
	Langugage string  `json:"language"`
	Books int 		  `json:"books"`
	Authors int 	  `json:"authors"`
	Fraction float64  `json:"fraction"`
}

type Status struct {
	Gutendexapi int  `json:"gutendexapi"`
	Langugaeapi int  `json:"languageapi"`
	Countriesapi int `json:"countriesapi"`
	Version string   `json:"version"`
	Uptime float64   `json:"uptime"`
}
