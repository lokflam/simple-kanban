package kanban

import (
	"fmt"
	"net/http"
	"unicode/utf8"

	"github.com/google/uuid"
	"github.com/lokflam/simple-kanban/internal/kanban/db"
	"github.com/lokflam/simple-kanban/internal/kanban/view"
)

type CardForm struct {
	ID      uuid.UUID
	Title   string
	Content string
}

func (c *CardForm) Bind(r *http.Request) error {
	if r == nil {
		return fmt.Errorf("request is nil")
	}

	idInput := r.FormValue("id")
	if idInput != "" {
		id, err := uuid.Parse(idInput)
		if err != nil {
			return &ErrInvalidID{err}
		}
		c.ID = id
	}

	c.Title = r.FormValue("title")
	c.Content = r.FormValue("content")
	return nil
}

func (c *CardForm) Validate() error {
	errs := make(map[string][]error)
	if c.Title == "" {
		errs["Title"] = append(errs["Title"], &ErrFieldRequired{"Title"})
	}
	if utf8.RuneCountInString(c.Content) > 5 {
		errs["Content"] = append(errs["Content"], &ErrFieldTooLong{"Content", 5})
	}
	if len(errs) > 0 {
		return &ErrInvalidFields{errs}
	}
	return nil
}

func CardViewModelFromCard(c db.Card) view.CardViewModel {
	return view.CardViewModel{
		ID:      c.ID,
		Title:   c.Title,
		Content: c.Content,
	}
}
