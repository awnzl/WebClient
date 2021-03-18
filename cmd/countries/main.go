package main

import (
	"fmt"
	"github.com/awnzl/lgTask1/internal/countryParser"
	"github.com/awnzl/lgTask1/internal/countryStorage"
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

	parser := countryParser.NewParser()
	parsedCountries, err := parser.Parse(file)
	if err != nil {
		fmt.Println(err)
	}

	countryStorage := countryStorage.NewStorage()
	countryStorage.InsertAll(parsedCountries)

	countryFinder := finder.NewFinder(countryStorage)
	foundCountries, err := countryFinder.Find(cfg.SearchOption, cfg.SearchArgument)

	if err != nil {
		fmt.Println(fmt.Sprintf("%v: %v", err, cfg.SearchArgument))
	}

	if foundCountries != nil {
		_publisher := publisher.NewStdoutPublisher()
		_publisher.Output(foundCountries)
	}
}
