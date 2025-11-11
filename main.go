package main

import (
	"backend/DBAdapter"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

const TEST_MODE bool = false

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	PORT := os.Getenv("PORT")
	USER := os.Getenv("USER")
	HOST := os.Getenv("HOST")
	DBPORT := os.Getenv("DBPORT")
	PASSWORD := os.Getenv("PASSWORD")
	DBNAME := os.Getenv("DBNAME")
	adapter := DBAdapter.InitAdapter(HOST, DBPORT, USER, PASSWORD, DBNAME)
	defer adapter.Close()

	if TEST_MODE {
		adapter.FillTestData()
	}

	engine := html.New("./html", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/greet", func(c *fiber.Ctx) error {
		return c.JSON("Hello pidor")
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("test", fiber.Map{})
	})
	app.Listen(":" + PORT)
}
