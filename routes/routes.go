package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/haile-paa/go-react/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/todos", controllers.GetTodos)
	app.Post("/api/todos", controllers.CreateTodo)
	app.Patch("/api/todos/:id", controllers.UpdateTodo)
	app.Delete("/api/todos/:id", controllers.DeleteTodo)
}
