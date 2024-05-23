package main

import (
	"encoding/json"
	"log"
	"net/http"
	"template-service/renderer"

	"github.com/gofiber/fiber/v3"
)

type data struct {
	Engine   string      `json:"engine"`
	Template string      `json:"template"`
	Model    interface{} `json:"model"`
}

func main() {

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {

		return c.SendString("This is a simple template render service. Use POST method to render templates.")
	})

	app.Post("/", func(c fiber.Ctx) error {
		bodyData := c.Body()

		var request data
		err := json.Unmarshal(bodyData, &request)
		if err != nil {
			log.Printf("Error: %s\n", err)
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		eng := renderer.GetRenderer(request.Engine)

		if eng == nil {
			log.Printf("Invalid Engine: %s\n", request.Engine)
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		out, err := eng.Render(request.Template, request.Model)
		if err != nil {
			log.Printf("Error: %s\n", err)
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		return c.SendString(out)

	})

	log.Fatal(app.Listen(":3100"))
}
