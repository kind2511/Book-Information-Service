package utilities

type GutendexResponse struct {
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Authors `json:"results"`
}

// Good
type BookCount struct {
    Count    int      `json:"count"`
}

// Good
type Authors struct {
    Authors      []struct {
        Name       string `json:"name"`
    } `json:"authors,omitempty"`
}

// Good
type TotalBookCount struct {
	TotalCount int  `json:"count"`
}

// Good
type Bookinfo struct {
	Language string  `json:"language"`
	Books int        `json:"books"`
	Authors int      `json:"authors"`
	Fraction float64 `json:"fraction"`
}

// Good
type CountryNameAndCode struct {
	Country string `json:"Official_Name"`
	Isocode string `json:"ISO3166_1_Alpha_2"`
}

// Good
type CountryPopulation struct {
	Readership int `json:"population"`
}

// Good
type CountryInfo struct {
	Country string `json:"country"`
	Isocode string `json:"isocode"`
	Books int      `json:"books"`
	Authors int    `json:"authors"`
	Readership int `json:"readership"`		
}

// Good
type Status struct {
	Gutendexapi int  `json:"gutendexapi"`
	Langugaeapi int  `json:"languageapi"`
	Countriesapi int `json:"countriesapi"`
	Version string   `json:"version"`
	Uptime float64   `json:"uptime"`
}
