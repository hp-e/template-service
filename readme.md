# Using the template service

This service is a simple template render service. It takes a template and a set of variables and renders the template with the variables.


## Usage

To use the service, you need to send a POST request to the service with the following JSON body:

```json
{
  "engine": "go",
  "template": "Hello, {{name}}!",
  "model": {
    "Name": "world"
  }
}
```

### Run with cURL

```bash
curl --location 'http://localhost:3100' \
--header 'Content-Type: application/json' \
--data '{
    "engine": "go",
    "template": "<p>{{ .Name}} is {{.Age}} years old, {{ if .HasCar}}and he has a car {{end}}. </p><ul>{{range .Languages}}<li>{{ . }}</li>{{end}}</ul>",
    "model": {
        "Name": "Ola",
        "Age": 30,
        "HasCar": true,
        "Languages": [
            "Go",
            "JavaScript"
        ]
    }
}'
```

```bash
curl --location 'http://localhost:3100' \
--header 'Content-Type: application/json' \
--data '{
    "engine": "django",
    "template": "<p>{{model.Name}} is {{model.Age}} years old, {% if model.HasCar %}and he has a car {% endif %}. </p><ul>{% for lang in model.Languages %}<li>{{ lang }}</li>{% endfor %}</ul>",
    "model": {
        "Name": "Ola",
        "Age": 30,
        "HasCar": true,
        "Languages": [
            "Go",
            "JavaScript"
        ]
    }
}'
```

## Model

This applies to the `model` field in the JSON body.
The model is a JSON object that contains the variables that will be used in the template. The keys in the model are the variable names and the values are the values that will be used in the template.
It must be a valid JSON object.
And it must start with a capital letter.

## Engine

This applies to the `engine` field in the JSON body.

The engine is either `go` or `django`. It specifies the template engine that will be used to render the template.

### Go

The Go template engine is a simple template engine that uses the `html/template` package in Go. It uses the `{{` and `}}` delimiters to denote variables in the template.

To access a property in the model, you must use the `.` operator. For example, to access the `Name` property in the model, you can use `{{.Name}}`.

Use https://golang.org/pkg/text/template/ for more information on the Go template language.

### Django

The Django template engine is a more powerful template engine that uses the Django template language. It uses the `{{` and `}}` delimiters to denote variables in the template.

To access a property in the Django model, you must use the `model.` operator. For example, to access the `Name` property in the model, you must use `{{model.Name}}`.

Use https://django.readthedocs.io/en/1.7.x/topics/templates.html for more information on the Django template language.