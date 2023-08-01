package main

import (
	"log"
	"todo-list/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/google/uuid"
)

func main() {
	db, err := repository.New()
	if err != nil {
		panic(err)
	}

	todos := make([]repository.Todo, 0)

	engine := html.New("./views", ".html")

	app := fiber.New(
		fiber.Config{
			Views:       engine,
			ViewsLayout: "layouts/main",
		},
	)
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Todos": todos,
		})
	})

	app.Post("/todo", func(c *fiber.Ctx) error {
		todo := repository.Todo{
			Id:     uuid.NewString(),
			Title:  c.FormValue("title"),
			IsDone: false,
		}

		todos = append(todos, todo)

		return c.Render("partials/item", todo)
	})

	app.Patch("/todo/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, v := range todos {
			if v.Id == id {
				log.Println(v.IsDone)
				todos[i].IsDone = !todos[i].IsDone

				return c.Render("partials/item", todos[i])
			}
		}

		return nil
	})

	app.Delete("/todo/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, v := range todos {
			if v.Id == id {
				todos = append(todos[:i], todos[i+1:]...)
				break
			}
		}

		log.Println(todos)

		return nil
	})

	app.Get("/layout", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("index", fiber.Map{
			"Todos": todos,
		}, "layouts/main")
	})

	app.Listen(":8080")
}
