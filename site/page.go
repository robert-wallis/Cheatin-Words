package site

import (
	"template"
	"os"
)

type Page struct {
	filename string
}

func (p *Page) Load(filename string) (t *template.Template, err os.Error) {
	p.filename = filename
	return template.ParseFile(p.filename, nil)
}
