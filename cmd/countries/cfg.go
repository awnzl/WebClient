package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/awnzl/lgTask1/internal/finder"
)

var ErrParseArgs = errors.New("cfg: incorrect argument")

type Config struct {
	SearchOption   finder.SearchOption
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

	if option == finder.SearchOptionUndefined {
		fmt.Println(ErrParseArgs, argument)
		flag.Usage()

		return Config{}, ErrParseArgs
	}

	return Config{
		SearchOption:   option,
		SearchArgument: argument,
	}, nil
}

func parseOptions(currencyCode, langCode, name string) (finder.SearchOption, string) {
	switch {
	case currencyCode != "":
		return finder.SearchOptionCurrencyCode, currencyCode
	case langCode != "":
		return finder.SearchOptionLang, langCode
	case name != "":
		return finder.SearchOptionName, name
	case len(flag.Args()) > 0:
		return finder.SearchOptionCountryCode, flag.Args()[0]
	default:
		return finder.SearchOptionUndefined, ""
	}
}

func usage() {
	fmt.Println("Usage of countries: [--currency-code=code|--lang-code=lang|--name=name] | country_code")
	flag.PrintDefaults()
}
