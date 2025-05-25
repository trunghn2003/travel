package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Place struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	ImagePath   string `json:"image_path"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Note        string `json:"note"`
	Suggestion  string `json:"suggestion"`
	CategoryID  int    `json:"category_id"`
	Area        string `json:"area"`
}

func main() {
	app := fiber.New()

	app.Use(cors.New())

	

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello World",
		})
	})

	app.Get("/travel", func(c *fiber.Ctx) error {
		// Read JSON file
		data, err := ioutil.ReadFile("data.json")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Could not read places data",
			})
		}

		var places []Place
		if err := json.Unmarshal(data, &places); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Could not parse places data",
			})
		}

		return c.JSON(places)
	})

	log.Fatal(app.Listen(":3000"))
}
