package finder

import (
	"errors"
	"github.com/awnzl/lgTask1/internal/countries"
	"strings"
)

type CountryFinder interface {
	Find(option, argument string, collection []countries.Country) ([]countries.Country, error)
	FindByCode(countryCode string, collection []countries.Country) ([]countries.Country, error)
	FindByCurrencyCode(currencyCode string, collection []countries.Country) ([]countries.Country, error)
	FindByLang(lang string, collection []countries.Country) ([]countries.Country, error)
	FindByName(name string, collection []countries.Country) ([]countries.Country, error)
}

type Finder struct { }

func NewFinder() Finder {
	return Finder{}
}

func (finder Finder) Find(option, argument string, collection []countries.Country) ([]countries.Country, error) {
	switch option {
	case "country-code":
		return finder.FindByCode(argument, collection)
	case "currency-code":
		return finder.FindByCurrencyCode(argument, collection)
	case "name":
		return finder.FindByName(argument, collection)
	case "lang":
		return finder.FindByLang(argument, collection)
	default:
		return nil, errors.New("can't find country with entered argument")
	}
}

func (finder Finder) FindByCode(countryCode string, collection []countries.Country) ([]countries.Country, error) {
	foundCountries := make([]countries.Country, 0)

	for _, country := range collection {
		if strings.ToLower(countryCode) == strings.ToLower(country.Alpha3Code) ||
			strings.ToLower(countryCode) == strings.ToLower(country.Alpha2Code) {
			foundCountries = append(foundCountries, country)
			break
		}
	}

	return finder.checkIfFound(foundCountries)
}

func (finder Finder) FindByCurrencyCode(currencyCode string, collection []countries.Country) ([]countries.Country, error) {
	foundCountries := make([]countries.Country, 0)

	for _, country := range collection {
		for _, currency := range country.Currencies {
			if strings.ToLower(currency.Code) == strings.ToLower(currencyCode) {
				foundCountries = append(foundCountries, country)
			}
		}
	}

	return finder.checkIfFound(foundCountries)
}

func (finder Finder) FindByLang(langIsoCode string, collection []countries.Country) ([]countries.Country, error) {
	foundCountries := make([]countries.Country, 0)

	for _, country := range collection {
		for _, language := range country.Languages {
			if strings.EqualFold(language.LanguageIsoCode, langIsoCode) {
				foundCountries = append(foundCountries, country)
			}
		}
	}

	return finder.checkIfFound(foundCountries)
}

func (finder Finder) FindByName(name string, collection []countries.Country) ([]countries.Country, error) {
	foundCountries := make([]countries.Country, 0)

	for _, country := range collection {
		if strings.Contains(strings.ToLower(country.Name), strings.ToLower(name)) {
			foundCountries = append(foundCountries, country)
		}
	}

	return finder.checkIfFound(foundCountries)
}

func (finder Finder) checkIfFound(foundCountries []countries.Country) ([]countries.Country, error) {
	switch {
	case len(foundCountries) == 0:
		return nil, errors.New("can't find country with entered argument")
	default:
		return foundCountries, nil
	}
}
