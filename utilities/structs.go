package utilities

type GutendexResponse struct {
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Authors `json:"results"`
}

type BookCount struct {
    Count    int      `json:"count"`
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
	Country string `json:"country"`
	Isocode string `json:"isocode"`
	Books int      `json:"books"`
	Authors int    `json:"authors"`
	Readership int `json:"readership"`		
}

type Status struct {
	Gutendexapi int  `json:"gutendexapi"`
	Langugaeapi int  `json:"languageapi"`
	Countriesapi int `json:"countriesapi"`
	Version string   `json:"version"`
	Uptime float64   `json:"uptime"`
}
