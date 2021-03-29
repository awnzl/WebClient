package webclient

import (
	"fmt"
	"net/http"

	"github.com/awnzl/lgTask1/internal/country"
	"github.com/awnzl/lgTask1/internal/parser"
)

const apiURI = "https://restcountries.eu/rest/v2"

type WebClient struct{}

func New() *WebClient {
	return &WebClient{}
}

func (c *WebClient) request(uri string) ([]country.Country, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response Status Code: %v", resp.StatusCode)
	}

	jsonParser := parser.New()
	countries, err := jsonParser.Parse(resp.Body)
	resp.Body.Close()

	if err != nil {
		return nil, err
	}

	return countries, nil
}

func (c *WebClient) ByCountryCode(code string) ([]country.Country, error) {
	requestURI := fmt.Sprintf("%v/alpha?codes=%v", apiURI, code)

	return c.request(requestURI)
}

func (c *WebClient) ByCurrencyCode(code string) ([]country.Country, error) {
	requestURI := fmt.Sprintf("%v/currency/%v", apiURI, code)

	return c.request(requestURI)
}

func (c *WebClient) ByLang(lang string) ([]country.Country, error) {
	requestURI := fmt.Sprintf("%v/lang/%v", apiURI, lang)

	return c.request(requestURI)
}

func (c *WebClient) ByName(name string) ([]country.Country, error) {
	requestURI := fmt.Sprintf("%v/name/%v?fullText=false", apiURI, name)

	return c.request(requestURI)
}
