package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

	_ "github.com/lib/pq"
)

// Database instance
var db *sql.DB

// Database settings
var (
	host     = os.Getenv("DB_HOST")
	port     = 5432 // Default port
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
)

// Connect function
func Connect() error {
	// set connection defaults
	if len(host) == 0 {
		host = "localhost"
	}
	if len(user) == 0 {
		user = "postgres"
	}
	if len(password) == 0 {
		password = "postgres"
	}
	if len(dbname) == 0 {
		dbname = "tech_demo"
	}

	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

func main() {

	// Connect with database
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋 with database connection!")
	})

	app.Listen(":3000")
}
