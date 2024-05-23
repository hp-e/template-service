package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"

	"github.com/gofiber/fiber/v3"
)

type data struct {
	Template string      `json:"template"`
	Model    interface{} `json:"model"`
}

func main() {

	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {

		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/", func(c fiber.Ctx) error {
		bodyData := c.Body()

		var request data
		err := json.Unmarshal(bodyData, &request)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		tmpl, err := template.New("test").Parse(request.Template)
		if err != nil {
			panic(err)
		}

		var doc bytes.Buffer

		err = tmpl.Execute(&doc, request.Model)
		if err != nil {
			panic(err)
		}

		return c.SendString(doc.String())
	})

	log.Fatal(app.Listen(":3100"))
}
