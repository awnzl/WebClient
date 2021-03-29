package countryservice

import (
	"fmt"
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

type WebClient interface {
	ByCountryCode(code string) ([]country.Country, error)
	ByCurrencyCode(code string) ([]country.Country, error)
	ByLang(lang string) ([]country.Country, error)
	ByName(name string) ([]country.Country, error)
}

type Writer interface {
	Write(countries []country.Country) error
}

type Service struct {
	client WebClient
	writer Writer
}

func New(client WebClient, writer Writer) *Service {
	return &Service{
		client: client,
		writer: writer,
	}
}

func (s *Service) Search(option SearchOption, arg string) error {
	countries, err := s.getCountries(option, arg)
	if err != nil {
		return fmt.Errorf("%w for %v %v", err, option, arg);
	}

	return s.writer.Write(countries)
}

func (s *Service) getCountries(option SearchOption, arg string) ([]country.Country, error) {
	switch option {
	case SearchOptionCurrencyCode:
		return s.client.ByCurrencyCode(arg)
	case SearchOptionName:
		return s.client.ByName(arg)
	case SearchOptionLang:
		return s.client.ByLang(arg)
	default:
		return s.client.ByCountryCode(arg)
	}
}
