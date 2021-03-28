package writer

import (
	"io"
	"text/template"

	"github.com/awnzl/lgTask1/internal/country"
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
	writer   io.Writer
}

func New(aWriter io.Writer) *Writer {
	writer := &Writer{
		Template: template.New("PrintCountry"),
		writer:   aWriter,
	}

	if _, err := writer.Template.Parse(outputTemplate); err != nil {
		panic("writer didn't parse template")
	}

	return writer
}

func (p *Writer) Write(countries []country.Country) error {
	for _, c := range countries {
		if err := p.Template.Execute(p.writer, c); err != nil {
			return err
		}
	}

	return nil
}
