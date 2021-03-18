package internal

import (
	"github.com/awnzl/lgTask1/internal/countries"
)

type Country = countries.Country

type CountryStorage interface {
	InsertAll(countries []Country) error
	//ByCountryCode(countryCode string) (Country, error)
	//ByCurrencyCode(countryCode string) (Country, error)
	//ByLanguageCode(countryCode string) (Country, error)
	//ByName(countryCode string) (Country, error)
}
