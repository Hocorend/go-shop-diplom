package main

import (
	"github.com/gofiber/fiber/v2"
	"go-shop-diplom/controllers"
	"go-shop-diplom/initializer"
	"go-shop-diplom/storage"
	"log"
	"os"
)

func Init() {
	initializer.LoadEnvVariables()
	storage.ConnectToDatabase()
	storage.MigrateModels()
}

func SetupRoutes(app *fiber.App) {
	api := app.Group("")
	api.Get("/products", controllers.GetProducts)
	api.Post("/users/deposit/add", controllers.AddDeposit)
	api.Post("/basket/add", controllers.AddToBasket)
	api.Post("/basket/pay", controllers.PayForBasket)
}

func main() {
	// Load environment variables
	Init()

	app := fiber.New()
	SetupRoutes(app)
	err := app.Listen(os.Getenv("SERVICE_HOST_PORT"))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
