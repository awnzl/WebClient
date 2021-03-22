package main

import (
	"fmt"
	"os"

	"github.com/awnzl/lgTask1/internal/finder"
	"github.com/awnzl/lgTask1/internal/parser"
	"github.com/awnzl/lgTask1/internal/writer"
)

func main() {
	cfg, err := ParseConfig()

	if err != nil {
		os.Exit(0)
	}

	const filename = "../countries.json"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Can't open file:", filename)
		fmt.Println("Error:", err)
		os.Exit(0)
	}

	parser := parser.New()
	parsedCountries, err := parser.Parse(file)
	if err != nil {
		fmt.Println("File parsing error:", filename)
		fmt.Println("Error:", err)
		os.Exit(0)
	}

	finder := finder.New()
	foundCountries := finder.Find(cfg.SearchOption, cfg.SearchArgument, parsedCountries)

	switch {
	case foundCountries != nil:
		writer := writer.New(os.Stdout)
		if err := writer.Write(foundCountries); err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println(fmt.Sprintf("can't find country with entered argument: %v", cfg.SearchArgument))
	}
}
