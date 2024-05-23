package renderer

type RenderService interface {
	Render(template string, model any) (string, error)
}

func GetRenderer(engine string) RenderService {
	if engine == "django" {
		return NewDjangoRenderer()
	} else if engine == "go" || engine == "" {
		return NewGoRenderer()
	}
	return nil
}
