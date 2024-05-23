package renderer

import "github.com/flosch/pongo2/v6"

type DjangoRenderer struct {
}

func NewDjangoRenderer() RenderService {
	return DjangoRenderer{}
}

func (d DjangoRenderer) Render(template string, model interface{}) (string, error) {
	tpl, err := pongo2.FromString(template)
	if err != nil {
		return "", err

	}

	ctx := pongo2.Context{}
	ctx["model"] = model
	out, err := tpl.Execute(ctx)
	if err != nil {
		return "", err
	}

	return out, nil
}
