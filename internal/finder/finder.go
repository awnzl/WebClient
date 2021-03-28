package finder

import (
	"strings"

	"github.com/awnzl/lgTask1/internal/country"
)

type SearchOption string

const (
	SearchOptionCountryCode  SearchOption = "country_code"
	SearchOptionCurrencyCode SearchOption = "currency_code"
	SearchOptionLang         SearchOption = "lang-code"
	SearchOptionName         SearchOption = "name"
	SearchOptionUndefined    SearchOption = "undefined"
)

func (o SearchOption) String() string {
	return string(o)
}

type Finder struct{}

func New() *Finder {
	return &Finder{}
}

func (finder *Finder) Find(option SearchOption, argument string, collection []country.Country) []country.Country {
	switch option {
	case SearchOptionCurrencyCode:
		return finder.FindByCurrencyCode(argument, collection)
	case SearchOptionName:
		return finder.FindByName(argument, collection)
	case SearchOptionLang:
		return finder.FindByLang(argument, collection)
	default:
		return finder.FindByCode(argument, collection)
	}
}

func (finder *Finder) FindByCode(countryCode string, collection []country.Country) []country.Country {
	var foundCountries []country.Country

	for _, country := range collection {
		if strings.EqualFold(countryCode, country.Alpha3Code) ||
			strings.EqualFold(countryCode, country.Alpha2Code) {

			foundCountries = append(foundCountries, country)
			break
		}
	}

	return foundCountries
}

func (finder *Finder) FindByCurrencyCode(currencyCode string, collection []country.Country) []country.Country {
	var foundCountries []country.Country

	for _, country := range collection {
		for _, currency := range country.Currencies {
			if strings.EqualFold(currency.Code, currencyCode) {
				foundCountries = append(foundCountries, country)
			}
		}
	}

	return foundCountries
}

func (finder *Finder) FindByLang(langIsoCode string, collection []country.Country) []country.Country {
	var foundCountries []country.Country

	for _, country := range collection {
		for _, language := range country.Languages {
			if strings.EqualFold(language.LanguageIsoCode, langIsoCode) {
				foundCountries = append(foundCountries, country)
			}
		}
	}

	return foundCountries
}

func (finder *Finder) FindByName(name string, collection []country.Country) []country.Country {
	var foundCountries []country.Country

	for _, country := range collection {
		if strings.Contains(strings.ToLower(country.Name), strings.ToLower(name)) {
			foundCountries = append(foundCountries, country)
		}
	}

	return foundCountries
}
