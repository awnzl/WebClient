package parser

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/awnzl/lgTask1/internal/countries"
)

type Parser struct { }

func New() *Parser {
	return &Parser{}
}

func (parser *Parser) Parse(r io.Reader) ([]countries.Country, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var countries = make([]countries.Country, 0)
	err = json.Unmarshal(data, &countries)

	return countries, err
}
