package main

import (
	"fmt"
	DB "go-learn/db"
	route "go-learn/routes"
	seeder "go-learn/seeders"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	DB.InitDB()
	// seed data
	fmt.Println("Seeding data...")
	seeder.SeedAll()
	fmt.Println("Seeding data success")

	app := fiber.New()
	// recover from panic
	app.Use(recover.New())
	route.InitRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
