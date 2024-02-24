package utilities

type BookInfoTemp struct {
    Count    int      `json:"count"`
	Results []Authors `json:"results"`
}

type Authors struct {
    Authors      []struct {
        Name       string `json:"name"`
    } `json:"authors,omitempty"`
}

type TotalBookCount struct {
	TotalCount int  `json:"count"`
}

type Bookinfo struct {
	Language string  `json:"language"`
	Books int        `json:"books"`
	Authors int      `json:"authors"`
	Fraction float64 `json:"fraction"`
}

type CountryNameAndCode struct {
	Country string `json:"Official_Name"`
	Isocode string `json:"ISO3166_1_Alpha_2"`
}

type CountryPopulation struct {
	Readership int `json:"population"`
}

type CountryInfo struct {
	Country string `json:"Official_Name"`
	Isocode string `json:"ISO3166_1_Alpha_2"`
	Books int      `json:"books"`
	Authors int    `json:"authors"`
	Readership int `json:"population"`		
}

type Status struct {
	Gutendexapi int  `json:"gutendexapi"`
	Langugaeapi int  `json:"languageapi"`
	Countriesapi int `json:"countriesapi"`
	Version string   `json:"version"`
	Uptime float64   `json:"uptime"`
}
