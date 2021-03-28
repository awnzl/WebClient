package main

import (
	"errors"
	"flag"
	"fmt"

	webclient "github.com/awnzl/lgTask1/internal/web_client"
)

var ErrParseArgs = errors.New("cfg: incorrect argument")

type Config struct {
	SearchOption   webclient.SearchOption
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

	if option == webclient.SearchOptionUndefined {
		fmt.Println(ErrParseArgs, argument)
		flag.Usage()

		return Config{}, ErrParseArgs
	}

	return Config{
		SearchOption:   option,
		SearchArgument: argument,
	}, nil
}

func parseOptions(currencyCode, langCode, name string) (webclient.SearchOption, string) {
	switch {
	case currencyCode != "":
		return webclient.SearchOptionCurrencyCode, currencyCode
	case langCode != "":
		return webclient.SearchOptionLang, langCode
	case name != "":
		return webclient.SearchOptionName, name
	case len(flag.Args()) > 0:
		return webclient.SearchOptionCountryCode, flag.Args()[0]
	default:
		return webclient.SearchOptionUndefined, ""
	}
}

func usage() {
	fmt.Println("Usage of countries: [--currency-code=code|--lang-code=lang|--name=name] | country_code")
	flag.PrintDefaults()
}
