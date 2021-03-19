package main

import (
	"fmt"
	"github.com/awnzl/lgTask1/internal/countryparser"
	"github.com/awnzl/lgTask1/internal/finder"
	"github.com/awnzl/lgTask1/internal/publisher"
	"os"
)

func main() {
	cfg := ParseConfig()

	file, err := os.OpenFile("../countries.json", os.O_RDONLY, 0400)
	if err != nil {
		fmt.Println(err)
	}

	parsedCountries, err := countryparser.NewParser().Parse(file)
	if err != nil {
		fmt.Println(err)
	}

	foundCountries, err := finder.NewFinder().Find(cfg.SearchOption, cfg.SearchArgument, parsedCountries)

	if err != nil {
		fmt.Println(fmt.Sprintf("%v: %v", err, cfg.SearchArgument))
	}

	if foundCountries != nil {
		publisher.NewStdoutPublisher().Output(foundCountries)
	}
}
