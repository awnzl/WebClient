package country

type CurrencyInfo struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type LanguageInfo struct {
	LanguageIsoCode string `json:"iso639_1"` // Iso639-1 Language Code
}

type Country struct {
	Name       string         `json:"name"`
	Alpha2Code string         `json:"alpha2Code"`
	Alpha3Code string         `json:"alpha3Code"`
	Capital    string         `json:"capital"`
	Region     string         `json:"region"`
	Subregion  string         `json:"subregion"`
	Population int64          `json:"population"`
	Currencies []CurrencyInfo `json:"currencies"`
	Languages  []LanguageInfo `json:"languages"`
}
