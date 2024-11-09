package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/haile-paa/go-react/config"
	"github.com/haile-paa/go-react/routes"
	"github.com/haile-paa/go-react/util"
)

func main() {
	fmt.Println("hello world")

	config.LoadEnv()

	dbClient := util.ConnectDB() // Ensures database connection setup
	defer dbClient.Disconnect(context.Background())
	fmt.Println("Connected to MongoDB")

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin,Content-Type,Accept",
	}))

	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	if os.Getenv("ENV") == "production" {
		app.Static("/", "./react-todo-front/dist")
	}
	log.Fatal(app.Listen("0.0.0.0:" + port))
}
