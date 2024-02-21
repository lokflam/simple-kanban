package view

import "github.com/google/uuid"

type CardViewModel struct {
	ID      uuid.UUID
	Title   string
	Content string
}

type CardFormViewModel struct {
	Open        bool
	Card        CardViewModel
	FieldErrors map[string][]error
}
