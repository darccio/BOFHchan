//go:generate pkger -o cmd/bofhchan
package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
	"github.com/imdario/BOFHchan/internal/post"
	"github.com/markbates/pkger"
)

var (
	mode = "debug"
)

func main() {
	app := fiber.New(&fiber.Settings{
		Views: views(),
	})
	post.RegisterURLs(app)
	if err := app.Listen(3000); err != nil {
		panic(err)
	}
}

func views() fiber.Views {
	dir := pkger.Dir("/templates")
	engine := html.NewFileSystem(dir, ".html")
	if mode == "debug" {
		engine.Debug(true)
		engine.Reload(true)
	}
	return engine
}
