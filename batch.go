package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ironarachne/namegen"
	"github.com/tjarratt/babble"
)

func DoBatchTask(name string) {
	// Connect with database
	if err := Connect(); err != nil {
		log.Fatalf("Failed to establish database connection: %v", err)
	}
	log.Println("Database connection established")

	switch name {
	case "create-todo":
		t := new(Todo)
		generator := namegen.NameGeneratorFromType("english", "both")
		name, err := generator.CompleteName("both")
		if err != nil {
			log.Fatalf("Failed to generate name: %v", err)
		}
		babbler := babble.NewBabbler()
		babbler.Separator = " "
		t.Description = fmt.Sprintf("%s %s", babbler.Babble(), name)
		t.DueDate = time.Now().AddDate(0, 0, 1)
		// Insert Todo into database
		_, err = db.Query("INSERT INTO todos (description, due_date) VALUES ($1, $2)", t.Description, t.DueDate)
		if err != nil {
			log.Fatalf("Failed to create todo: %v", err)
		}
		// Print result
		log.Printf("Created %s todo item\n", name)
	}
}
