package main

import (
	"fmt"
	"os"
	"strconv"

	_ "database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"backend/DBAdapter"
	"backend/DBConnection"
)

type BaseTpuUser struct {
	uuid        uuid.UUID
	name        string
	last_name   string
	midlde_name string
	login       string
}

func init_adapter() *DBAdapter.DBAdapter {
	port, err := strconv.Atoi(os.Getenv("DBPORT"))
	if err != nil {
		fmt.Println("DBPort is not a number")
	}
	user := os.Getenv("USER")
	host := os.Getenv("HOST")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	connInfo := DBConnection.ConnectionInfo{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbname,
	}

	adapter, err := DBAdapter.InitAdapter(connInfo)
	if err != nil {
		panic(err)
	}
	return adapter
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	adapter := init_adapter()
	port := os.Getenv("PORT")

	engine := html.New("./html", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	fmt.Println("Hello backend")
	app.Get("/greet", func(c *fiber.Ctx) error {
		return c.JSON("Hello pidor")
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("test", fiber.Map{})
	})
	app.Listen(":" + port)
}
