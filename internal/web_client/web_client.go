package webclient

import (
	"fmt"
	"io"
	"net/http"
)

const apiURI = "https://restcountries.eu/rest/v2"

type SearchOption string

const (
	SearchOptionCountryCode  SearchOption = "country_code"
	SearchOptionCurrencyCode SearchOption = "currency_code"
	SearchOptionLang         SearchOption = "lang-code"
	SearchOptionName         SearchOption = "name"
	SearchOptionUndefined    SearchOption = "undefined"
)

func (o SearchOption) String() string {
	return string(o)
}

type WebClient struct{}

func New() *WebClient {
	return &WebClient{}
}

func (c *WebClient) Get(option SearchOption, argument string) (io.ReadCloser, error) {
	var requestURI string

	switch option {
	case SearchOptionCurrencyCode:
		requestURI = fmt.Sprintf("%v/currency/%v", apiURI, argument)
	case SearchOptionName:
		requestURI = fmt.Sprintf("%v/name/%v?fullText=false", apiURI, argument)
	case SearchOptionLang:
		requestURI = fmt.Sprintf("%v/lang/%v", apiURI, argument)
	default:
		requestURI = fmt.Sprintf("%v/alpha?codes=%v", apiURI, argument)
	}

	reader, err := c.request(requestURI)
	if err != nil {
		return nil, fmt.Errorf("%v\n\tsearch option: %v\n\targument: %v", err, option, argument)
	}

	return reader, err
}

func (c *WebClient) request(uri string) (io.ReadCloser, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response Status Code: %v", resp.StatusCode)
	}

	return resp.Body, nil
}
