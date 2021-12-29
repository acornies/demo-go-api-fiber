package main

import "time"

type Todo struct {
	ID          int       `json:id`
	Description string    `json:description`
	DueDate     time.Time `json:due_date`
}

type Todos struct {
	Todos []Todo `json:todos`
}
