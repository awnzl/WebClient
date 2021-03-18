package countryStorage

import (
	"github.com/awnzl/lgTask1/internal/countries"
)

type Country = countries.Country

type Storage struct {
	Countries []Country
}

func NewStorage() Storage {
	return Storage{make([]Country, 0)}
}

func (s *Storage) InsertAll(countries []Country) error {
	s.Countries = countries
	return nil
}

//func (s *Storage) ByCountryCode(countryCode string) (Country, error) {
//
//}
//
//func (s *Storage) ByCurrencyCode(currencyCode string) (Country, error) {
//
//}
//
//func (s *Storage) ByLanguageCode(languageIsoCode string) (Country, error) {
//
//}
//
//func (s *Storage) ByName(name string) (Country, error) {
//
//}
