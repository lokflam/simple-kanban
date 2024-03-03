package kanban

import "github.com/google/uuid"

type Status struct {
	ID   uuid.UUID
	Name string
}

type Card struct {
	ID      uuid.UUID
	Title   string
	Content string
	Status  Status
}
