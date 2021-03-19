package countryparser

import (
	"encoding/json"
	"github.com/awnzl/lgTask1/internal/countries"
	"io"
	"io/ioutil"
)

type CountriesParser interface {
	Parse(r io.Reader) ([]countries.Country, error)
}

type Parser struct { }

func NewParser() Parser {
	return Parser{}
}

func (parser Parser) Parse(r io.Reader) ([]countries.Country, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var countries = make([]countries.Country,	0)
	unmarshalErr := json.Unmarshal(data, &countries)

	return countries, unmarshalErr
}
