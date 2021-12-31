package main

import (
	"database/sql"
	"flag"
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
	mode     = flag.String("mode", "batch", "The binary execution mode (batch | server)")
	task     = flag.String("task", "", "The task to perform in batch mode")
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

	flag.Parse()

	switch *mode {
	case "batch":
		DoBatchTask(*task)
	case "server":
		// Connect with database
		if err := Connect(); err != nil {
			log.Fatalf("Failed to establish database connection: %v", err)
		}
		log.Println("Database connection established")
		app := fiber.New()

		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello, World 👋 with database connection!")
		})

		app.Get("/todos", func(c *fiber.Ctx) error {
			// Select all todo items from database
			rows, err := db.Query("SELECT id, description, due_date FROM todos order by id DESC")
			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
			defer rows.Close()
			result := Todos{}

			for rows.Next() {
				todo := Todo{}
				if err := rows.Scan(&todo.ID, &todo.Description, &todo.DueDate); err != nil {
					return err // Exit if we get an error
				}

				// Append Employee to Employees
				result.Todos = append(result.Todos, todo)
			}
			// Return Employees in JSON format
			return c.JSON(result)
		})

		app.Listen(":3000")
	default:
		log.Fatalf("Unsupported execution mode %s. Must be \"batch\" or \"server\".\n", *mode)
	}
}
