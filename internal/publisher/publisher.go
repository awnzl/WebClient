package publisher

import (
	"github.com/awnzl/lgTask1/internal/countries"
	"os"
	"text/template"
)

type Publisher interface {
	Output([]countries.Country)
}

type StdoutPublisher struct {
	Template *template.Template
}

func NewStdoutPublisher() *StdoutPublisher {
	publisher := &StdoutPublisher {
		template.New("PrintCountry"),
	}

	templatePattern := `name: {{ .Name }}
capital: {{ .Capital }}
region: {{ .Region }}
subregion: {{ .Subregion }}
population: {{ .Population }}
languages: {{ range $v := .Languages }}
 - {{$v.LanguageIsoCode}}{{ end }}
currencies: {{ range $v := .Currencies }}
 - code: {{$v.Code}}
   name: {{$v.Name}}
   symbol: {{$v.Symbol}}{{ end }}

`
	if _, err := publisher.Template.Parse(templatePattern); err != nil {
		panic("publisher didn't parse template")
	}

	return publisher
}

func (p *StdoutPublisher) Output(aCountries []countries.Country) {
	for _, country := range aCountries {
		p.Template.Execute(os.Stdout, country)
	}
}
