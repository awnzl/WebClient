package main

import (
	"fmt"
	"os"

	"github.com/awnzl/lgTask1/internal/parser"
	webclient "github.com/awnzl/lgTask1/internal/web_client"
	"github.com/awnzl/lgTask1/internal/writer"
)

func main() {
	cfg, err := parseConfig()
	if err != nil {
		os.Exit(0)
	}

	webClient := webclient.New()
	jsonParser := parser.New()
	outputWriter := writer.New(os.Stdout)

	data, err := webClient.Get(cfg.SearchOption, cfg.SearchArgument)
	if err != nil {
		fmt.Println("Web data request failure")
		fmt.Println("Error:", err)
		data.Close()
		os.Exit(0)
	}

	countries, err := jsonParser.Parse(data)

	data.Close()

	if err != nil {
		fmt.Println("Data parsing error")
		fmt.Println("Error:", err)
		os.Exit(0)
	}

	if len(countries) == 0 {
		fmt.Printf("can't find country with entered argument: %v\n", cfg.SearchArgument)
		os.Exit(0)
	}

	if err := outputWriter.Write(countries); err != nil {
		fmt.Println(err)
	}
}
