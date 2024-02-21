package utilities

type BookInfoTemp struct {
    Count    int        `json:"count"`
	Authors  []string 	  `json:"authors"`
}

type TotalBookCount struct {
	TotalCount int `json:"count"`
}

type Bookinfo struct {
	Language string  `json:"language"`
	Books int 		  `json:"books"`
	Authors int 	  `json:"authors"`
	Fraction float64  `json:"fraction"`
}

//Status Endpoint:
//-----------------------------------------------------

type Status struct {
	Gutendexapi int  `json:"gutendexapi"`
	Langugaeapi int  `json:"languageapi"`
	Countriesapi int `json:"countriesapi"`
	Version string   `json:"version"`
	Uptime float64   `json:"uptime"`
}
