package main

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	SearchOption string
	SearchArgument string
}

func ParseConfig() Config {
	defineFlags()
	flag.Parse()

	searchOption, searchArgument := "", ""
	flag.Visit(func(f *flag.Flag) {
		if f.Value.String() != "" {
			searchOption = f.Name
			searchArgument = f.Value.String()
		}
	})

	if searchOption == "" && len(flag.Args()) > 0 {
		searchOption = "country-code"
		searchArgument = flag.Args()[0]
	}

	if searchArgument == "" {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n  .%s country_code|--flag=value\n",
			os.Args[0], os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}

	return Config{searchOption, searchArgument}
}

func defineFlags() {
	flag.String("currency-code", "", "--currency-code=cad")
	flag.String("name", "", "--name=Canada")
	flag.String("lang", "", "--lang-code=en")
}