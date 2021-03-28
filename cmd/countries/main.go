package main

import (
	"fmt"
	"os"

	"github.com/awnzl/lgTask1/internal/finder"
	"github.com/awnzl/lgTask1/internal/parser"
	"github.com/awnzl/lgTask1/internal/writer"
)

const filename = "../countries.json"

func main() {
	cfg, err := parseConfig()
	if err != nil {
		os.Exit(0)
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Can't open file:", filename)
		fmt.Println("Error:", err)
		os.Exit(0)
	}

	countriesFinder := finder.New()
	jsonParser := parser.New()
	outputWriter := writer.New(os.Stdout)

	parsedCountries, err := jsonParser.Parse(file)
	if err != nil {
		fmt.Println("File parsing error:", filename)
		fmt.Println("Error:", err)
		os.Exit(0)
	}

	foundCountries := countriesFinder.Find(cfg.SearchOption, cfg.SearchArgument, parsedCountries)

	if len(foundCountries) == 0 {
		fmt.Printf("can't find country with entered argument: %v\n", cfg.SearchArgument)
		os.Exit(0)
	}

	if err := outputWriter.Write(foundCountries); err != nil {
		fmt.Println(err)
	}
}
