package main

import (
	"context"
	"database/sql"
	"todo-list/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

func main() {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		panic(err)
	}

	// create tables
	if _, err := db.ExecContext(context.Background(), ddl); err != nil {
		panic(err)
	}

	repository := repository.New(db)
	engine := html.New("./views", ".html")

	app := fiber.New(
		fiber.Config{
			Views:       engine,
			ViewsLayout: "layouts/main",
		},
	)
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		todos, err := repository.GetAllTodos()
		if err != nil {
			return err
		}

		return c.Render("index", fiber.Map{
			"Todos": todos,
		})
	})

	app.Post("/todo", func(c *fiber.Ctx) error {
		todo, err := repository.CreateTodo(c.FormValue("title"))
		if err != nil {
			return err
		}

		return c.Render("partials/item", todo)
	})

	app.Patch("/todo/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		todo, err := repository.UpdateStatus(id)
		if err != nil {
			return err
		}

		return c.Render("partials/item", todo)
	})

	app.Delete("/todo/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		return repository.DeleteById(id)
	})

	app.Get("/layout", func(c *fiber.Ctx) error {
		todos, err := repository.GetAllTodos()
		if err != nil {
			return err
		}

		// Render index within layouts/main
		return c.Render("index", fiber.Map{
			"Todos": todos,
		}, "layouts/main")
	})

	app.Listen(":8080")
}
