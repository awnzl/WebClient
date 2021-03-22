package writer

import (
	"github.com/awnzl/lgTask1/internal/countries"
	"io"
	"text/template"
)

const outputTemplate = `name: {{ .Name }}
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

type Writer struct {
	Template *template.Template
	writer io.Writer
}

func New(aWriter io.Writer) *Writer {
	writer := &Writer{
		template.New("PrintCountry"),
		aWriter,
	}

	if _, err := writer.Template.Parse(outputTemplate); err != nil {
		panic("writer didn't parse template")
	}

	return writer
}

func (p *Writer) Write(aCountries []countries.Country) error {
	for _, country := range aCountries {
		if err := p.Template.Execute(p.writer, country); err != nil {
			return err
		}
	}
	return nil
}
