package post

import (
	"github.com/gofiber/fiber"
)

func RegisterURLs(app *fiber.App) {
	router := app.Group("/posts")
	router.Get("/:id", get)
	router.Post("/", post)
}

func get(c *fiber.Ctx) {
	err := c.Render("post/post", fiber.Map{}, "layout")
	if err != nil {
		c.Next(err)
	}
}

func post(c *fiber.Ctx) {

}
