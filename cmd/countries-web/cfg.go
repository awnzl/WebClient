package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/awnzl/lgTask1/internal/countryservice"
)

var ErrParseArgs = errors.New("cfg: incorrect argument")

type Config struct {
	SearchOption   countryservice.SearchOption
	SearchArgument string
}

func parseConfig() (Config, error) {
	var currencyCode, langCode, name string

	flag.StringVar(&currencyCode, "currency-code", "", "--currency-code=cad")
	flag.StringVar(&langCode, "lang-code", "", "--lang-code=en")
	flag.StringVar(&name, "name", "", "--name=Canada")

	flag.Usage = usage
	flag.Parse()

	option, argument := parseOptions(currencyCode, langCode, name)

	if option == countryservice.SearchOptionUndefined {
		fmt.Println(ErrParseArgs, argument)
		flag.Usage()

		return Config{}, ErrParseArgs
	}

	return Config{
		SearchOption:   option,
		SearchArgument: argument,
	}, nil
}

func parseOptions(currencyCode, langCode, name string) (countryservice.SearchOption, string) {
	switch {
	case currencyCode != "":
		return countryservice.SearchOptionCurrencyCode, currencyCode
	case langCode != "":
		return countryservice.SearchOptionLang, langCode
	case name != "":
		return countryservice.SearchOptionName, name
	case len(flag.Args()) > 0:
		return countryservice.SearchOptionCountryCode, flag.Args()[0]
	default:
		return countryservice.SearchOptionUndefined, ""
	}
}

func usage() {
	fmt.Println("Usage of countries: [--currency-code=code|--lang-code=lang|--name=name] | country_code")
	flag.PrintDefaults()
}
