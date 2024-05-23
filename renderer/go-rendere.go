package renderer

import (
	"bytes"
	tpl "html/template"
)

type GoRenderer struct{}

func NewGoRenderer() RenderService {
	return GoRenderer{}
}

func (g GoRenderer) Render(template string, model interface{}) (string, error) {
	tmpl, err := tpl.New("test").Parse(template)
	if err != nil {
		return "", err
	}

	var doc bytes.Buffer

	err = tmpl.Execute(&doc, model)
	if err != nil {
		return "", err

	}
	return doc.String(), nil
}
