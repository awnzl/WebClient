package main

import (
	"fmt"
	"os"

	"github.com/awnzl/lgTask1/internal/countryservice"
	"github.com/awnzl/lgTask1/internal/webclient"
	"github.com/awnzl/lgTask1/internal/writer"
)

func main() {
	cfg, err := parseConfig()
	if err != nil {
		os.Exit(0)
	}

	service := countryservice.New(webclient.New(), writer.New(os.Stdout))

	if err := service.Search(cfg.SearchOption, cfg.SearchArgument); err != nil {
		fmt.Println(err)
	}
}
