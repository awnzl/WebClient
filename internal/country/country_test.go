package country

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const countryJSON = `
{
    "name": "United States of America",
    "topLevelDomain": [
      ".us"
    ],
    "alpha2Code": "US",
    "alpha3Code": "USA",
    "callingCodes": [
      "1"
    ],
    "capital": "Washington, D.C.",
    "altSpellings": [
      "US",
      "USA",
      "United States of America"
    ],
    "region": "Americas",
    "subregion": "Northern America",
    "population": 323947000,
    "latlng": [
      38,
      -97
    ],
    "demonym": "American",
    "area": 9629091,
    "gini": 48,
    "timezones": [
      "UTC-12:00",
      "UTC-11:00",
      "UTC-10:00",
      "UTC-09:00",
      "UTC-08:00",
      "UTC-07:00",
      "UTC-06:00",
      "UTC-05:00",
      "UTC-04:00",
      "UTC+10:00",
      "UTC+12:00"
    ],
    "borders": [
      "CAN",
      "MEX"
    ],
    "nativeName": "United States",
    "numericCode": "840",
    "currencies": [
      {
        "code": "USD",
        "name": "United States dollar",
        "symbol": "$"
      }
    ],
    "languages": [
      {
        "iso639_1": "en",
        "iso639_2": "eng",
        "name": "English",
        "nativeName": "English"
      }
    ],
    "translations": {
      "de": "Vereinigte Staaten von Amerika",
      "es": "Estados Unidos",
      "fr": "États-Unis",
      "ja": "アメリカ合衆国",
      "it": "Stati Uniti D'America",
      "br": "Estados Unidos",
      "pt": "Estados Unidos",
      "nl": "Verenigde Staten",
      "hr": "Sjedinjene Američke Države",
      "fa": "ایالات متحده آمریکا"
    },
    "flag": "https://restcountries.eu/data/usa.svg",
    "regionalBlocs": [
      {
        "acronym": "NAFTA",
        "name": "North American Free Trade Agreement",
        "otherAcronyms": [],
        "otherNames": [
          "Tratado de Libre Comercio de América del Norte",
          "Accord de Libre-échange Nord-Américain"
        ]
      }
    ],
    "cioc": "USA"
}
`

func TestCountryUnmarshaling(t *testing.T) {
	want := Country{
		Name:       "United States of America",
		Alpha2Code: "US",
		Alpha3Code: "USA",
		Capital:    "Washington, D.C.",
		Region:     "Americas",
		Subregion:  "Northern America",
		Population: 323947000,
		Languages: []LanguageInfo{
			LanguageInfo{LanguageIsoCode: "en"},
		},
		Currencies: []CurrencyInfo{
			CurrencyInfo{
				Code:   "USD",
				Name:   "United States dollar",
				Symbol: "$",
			},
		},
	}

	var got Country
	if err := json.Unmarshal([]byte(countryJSON), &got); err != nil {
		t.Errorf("some error occured during unmarshalling")
	}

	assert.Equal(t, want, got, "Unmarshalling Country fails")
}
