package finder

import (
	"errors"
	"github.com/awnzl/lgTask1/internal/countries"
	"github.com/awnzl/lgTask1/internal/countryStorage"
	"strings"
)

type Country = countries.Country

type CountryFinder interface {
	Find(option, argument string) ([]Country, error)
	FindByCode(countryCode string) ([]Country, error)
	FindByCurrencyCode(currencyCode string) ([]Country, error)
	FindByLang(lang string) ([]Country, error)
	FindByName(name string) ([]Country, error)
}

type Finder struct {
	Storage countryStorage.Storage
}

func NewFinder(storage countryStorage.Storage) Finder {
	return Finder{storage}
}

func (finder *Finder) Find(option, argument string) ([]Country, error) {
	switch option {
	case "country-code":
		return finder.FindByCode(argument)
	case "currency-code":
		return finder.FindByCurrencyCode(argument)
	case "name":
		return finder.FindByName(argument)
	case "lang":
		return finder.FindByLang(argument)
	default:
		return nil, errors.New("can't find country with entered argument")
	}
}

func (finder *Finder) FindByCode(countryCode string) ([]Country, error) {
	_countries := make([]Country, 0)

	for _, country := range finder.Storage.Countries {
		if strings.ToLower(countryCode) == strings.ToLower(country.Alpha3Code) ||
			strings.ToLower(countryCode) == strings.ToLower(country.Alpha2Code) {
			_countries = append(_countries, country)
			break
		}
	}

	return finder.checkIfFound(_countries)
}

func (finder *Finder) FindByCurrencyCode(currencyCode string) ([]Country, error) {
	_countries := make([]Country, 0)

	for _, country := range finder.Storage.Countries {
		for _, currency := range country.Currencies {
			if strings.ToLower(currency.Code) == strings.ToLower(currencyCode) {
				_countries = append(_countries, country)
			}
		}
	}

	return finder.checkIfFound(_countries)
}

func (finder *Finder) FindByLang(langIsoCode string) ([]Country, error) {
	_countries := make([]Country, 0)

	for _, country := range finder.Storage.Countries {
		for _, language := range country.Languages {
			if strings.EqualFold(language.LanguageIsoCode, langIsoCode) {
				_countries = append(_countries, country)
			}
		}
	}

	return finder.checkIfFound(_countries)
}

func (finder *Finder) FindByName(name string) ([]Country, error) {
	_countries := make([]Country, 0)

	for _, country := range finder.Storage.Countries {
		if strings.Contains(strings.ToLower(country.Name), strings.ToLower(name)) {
			_countries = append(_countries, country)
		}
	}

	return finder.checkIfFound(_countries)
}

func (finder *Finder) checkIfFound(_countries []Country) ([]Country, error) {
	switch {
	case len(_countries) == 0:
		return nil, errors.New("can't find country with entered argument")
	default:
		return _countries, nil
	}
}