package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/awnzl/lgTask1/internal/finder"
)

var ParseArgsError = errors.New("cfg: incorrect argument")

type Config struct {
	SearchOption finder.SearchOption
	SearchArgument string
}

func ParseConfig() (Config, error) {
	defineFlags()
	flag.Usage = usage
	flag.Parse()

	option, argument := finder.SearchOptionUndefined, ""
	flag.Visit(func(f *flag.Flag) {
		switch f.Value.String() != "" {
		case f.Name == "currency-code":
			option, argument = finder.SearchOptionCurrencyCode, f.Value.String()
		case f.Name == "lang-code":
			option, argument = finder.SearchOptionLang, f.Value.String()
		case f.Name == "name":
			option, argument = finder.SearchOptionName, f.Value.String()
		}
	})

	if option == finder.SearchOptionUndefined && len(flag.Args()) > 0 {
		option, argument = finder.SearchOptionCountryCode, flag.Args()[0]
	}

	if option == finder.SearchOptionUndefined {
		fmt.Println(ParseArgsError, argument)
		flag.Usage()
		return Config{}, ParseArgsError
	}

	return Config{
		SearchOption: option,
		SearchArgument: argument,
	}, nil
}

func defineFlags() {
	flag.String("currency-code", "", "--currency-code=cad")
	flag.String("lang-code", "", "--lang-code=en")
	flag.String("name", "", "--name=Canada")
}

func usage() {
	fmt.Println("Usage of countries: [--currency-code=code|--lang-code=lang|--name=name] | country_code")
	flag.PrintDefaults()
}