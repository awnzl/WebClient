package countryParser

import (
	"encoding/json"
	"github.com/awnzl/lgTask1/internal/countries"
	"io"
	"io/ioutil"
)

type CountriesParser interface {
	Parse(r io.Reader) ([]Country, error)
}

type Country = countries.Country

type Parser struct {
}

func NewParser() Parser {
	return Parser{}
}

func (parser *Parser) Parse(r io.Reader) ([]Country, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var countries = make([]Country,	0)
	json.Unmarshal(data, &countries)

	return countries, nil
}
