package post

import (
	"github.com/gofiber/fiber"
)

func RegisterURLs(app *fiber.App) {
	router := app.Group("/posts")
	router.Get("/submit", publish)
	router.Get("/:id", get)
	router.Post("/", post)
}

func get(c *fiber.Ctx) {
	context := fiber.Map{
		"id": c.Params("id"),
	}
	err := c.Render("post/post", context, "layout")
	if err != nil {
		c.Next(err)
	}
}

func publish(c *fiber.Ctx) {
	err := c.Render("post/publish", fiber.Map{}, "layout")
	if err != nil {
		c.Next(err)
	}
}

func post(c *fiber.Ctx) {

}
