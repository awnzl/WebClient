package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

var foundFlag = false

type CurrencyInfo struct {
	Code string				`json:"code"`
	Name string				`json:"name"`
	Symbol string			`json:"symbol"`
}

type LanguageInfo struct {
	Iso639_1 string			`json:"iso639_1"`
}

type Country struct {
	Name string               `json:"name"`
	Alpha2Code string         `json:alpha2Code`
	Alpha3Code string         `json:"alpha3Code"`
	Capital string            `json:"capital"`
	Region string             `json:"region"`
	Subregion string          `json:"subregion"`
	Population int64          `json:"population"`
	Currencies []CurrencyInfo `json:"currencies"`
	Languages []LanguageInfo  `json:"languages"`
}

func main() {
	defineFlags()
	flag.Parse()

	if len(os.Args) == 1 {
		func() {
			fmt.Fprintf(os.Stderr, "Usage of %s:\n  .%s country_code|--flag=value\n", os.Args[0], os.Args[0])
			flag.PrintDefaults()
			os.Exit(0)
		}()
	}

	if foundFlag {
		os.Exit(0)
	}

	if len(flag.Args()) > 0 {
		findByCode(flag.Args()[0])
	}

	if !foundFlag {
		fmt.Println("Can't find country with entered argument:", os.Args[1])
	}
}

func findByName(name string) error {
	countries := getJSON()

	for _, country := range countries {
		if len(name) > 0 &&
			strings.Contains(strings.ToLower(country.Name), strings.ToLower(name)) {
			foundFlag = true
			printCountry(&country)
			fmt.Println()
		}
	}

	return nil
}

func findByLang(lang string) error {
	countries := getJSON()

	for _, country := range countries {
		for _, language := range country.Languages {
			if strings.ToLower(language.Iso639_1) == strings.ToLower(lang) {
				foundFlag = true
				printCountry(&country)
				fmt.Println()
			}
		}
	}

	return nil
}

func findByCurrencyCode(currencyCode string) error {
	countries := getJSON()

	for _, country := range countries {
		for _, currency := range country.Currencies {
			if strings.ToLower(currency.Code) == strings.ToLower(currencyCode) {
				foundFlag = true
				printCountry(&country)
				fmt.Println()
			}
		}
	}

	return nil
}

func findByCode(countryCode string) {
	countries := getJSON()

	for _, country := range countries {
		if strings.ToLower(countryCode) == strings.ToLower(country.Alpha3Code) ||
			strings.ToLower(countryCode) == strings.ToLower(country.Alpha2Code) {
			printCountry(&country)
			foundFlag = true
			break
		}
	}
}

func printCountry(country *Country) {
	template := template.New("PrintCountry")
	templatePattern :=
`name {{ .Name }}
capital: {{ .Capital }}
region: {{ .Region }}
subregion: {{ .Subregion }}
population: {{ .Population }}
languages: {{ range $v := .Languages }}
 - {{$v.Iso639_1}}{{ end }}
currencies: {{ range $v := .Currencies }}
 - code: {{$v.Code}}
   name: {{$v.Name}}
   symbol: {{$v.Symbol}}{{ end }}
`

	template.Parse(templatePattern)
	template.Execute(os.Stdout, country)
}

func getJSON() []Country {
	file, err := os.OpenFile("./countries.json", os.O_RDONLY, 0400)
	if err != nil {
		fmt.Println("oh shit!", err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("no way!", err)
	}

	var countries = make([]Country,	0)
	json.Unmarshal(data, &countries)

	return countries
}

func defineFlags() {
	flag.Func("currency-code", "Country currency code", findByCurrencyCode)
	flag.Func("name", "Country name", findByName)
	flag.Func("lang", "Country language", findByLang)
}
