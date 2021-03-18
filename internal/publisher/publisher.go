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
}

func NewStdoutPublisher() StdoutPublisher {
	return StdoutPublisher{}
}

func (p *StdoutPublisher) Output(countires []countries.Country) {
	template := template.New("PrintCountry")
	templatePattern :=
		`name: {{ .Name }}
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
	template.Parse(templatePattern)

	for _, country := range countires {
		template.Execute(os.Stdout, country)
	}
}