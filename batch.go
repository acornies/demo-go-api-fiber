package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/ironarachne/namegen"
)

var origins = []string{
	"anglosaxon",
	"dutch",
	"dwarf",
	"elf",
	"english",
	"estonian",
	"fantasy",
	"finnish",
	"french",
	"german",
	"greek",
	"hindu",
	"icelandic",
	"indonesian",
	"italian",
	"japanese",
	"korean",
	"mayan",
	"nepalese",
	"norwegian",
	"portuguese",
	"russian",
	"spanish",
	"swedish",
	"thai",
}

func DoBatchTask(name string) {
	// Connect with database
	if err := Connect(); err != nil {
		log.Fatalf("Failed to establish database connection: %v", err)
	}
	log.Println("Database connection established")

	switch name {
	case "create-todo":
		t := new(Todo)
		rand.Seed(time.Now().Unix())
		pick := origins[rand.Intn(len(origins))]
		generator := namegen.NameGeneratorFromType(pick, "both")
		fullName, err := generator.CompleteName("both")
		if err != nil {
			log.Fatalf("Failed to generate full name: %v", err)
		}
		t.Description = fullName
		t.DueDate = time.Now().AddDate(0, 0, 1)
		// Insert Todo into database
		_, err = db.Query("INSERT INTO todos (description, due_date) VALUES ($1, $2)", t.Description, t.DueDate)
		if err != nil {
			log.Fatalf("Failed to create todo: %v", err)
		}
		// Print result
		log.Printf("Created %s todo item\n", fullName)
	}
}
