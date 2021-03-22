package main

import (
	"fmt"
	"github.com/awnzl/lgTask1/internal/finder"
	"github.com/awnzl/lgTask1/internal/parser"
	"github.com/awnzl/lgTask1/internal/writer"
	"os"
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

	parsedCountries, err := parser.New().Parse(file)
	if err != nil {
		fmt.Println("File parsing error:", filename)
		fmt.Println("Error:", err)
		os.Exit(0)
	}

	foundCountries := finder.New().Find(cfg.SearchOption, cfg.SearchArgument, parsedCountries)

	switch {
	case foundCountries != nil:
		if err := writer.New(os.Stdout).Write(foundCountries); err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println(fmt.Sprintf("can't find country with entered argument: %v", cfg.SearchArgument))
	}
}
