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
	ID       uuid.UUID
	Title    string
	Content  string
	StatusID uuid.UUID
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

	statusIDInput := r.FormValue("status_id")
	if statusIDInput != "" {
		statusID, err := uuid.Parse(statusIDInput)
		if err != nil {
			return &ErrInvalidID{err}
		}
		c.StatusID = statusID
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
	var id string
	if c.ID != uuid.Nil {
		id = c.ID.String()
	}

	var statusID string
	if c.StatusID != uuid.Nil {
		statusID = c.StatusID.String()
	}

	return view.CardViewModel{
		ID:       id,
		StatusID: statusID,
		Title:    c.Title,
		Content:  c.Content,
	}
}

func StatusViewModelsFromListStatusesRows(statuses []db.StatusesRow) []view.StatusViewModel {
	viewModels := make([]view.StatusViewModel, 0, len(statuses))
	for _, s := range statuses {
		viewModels = append(viewModels, StatusViewModelFromStatus(s))
	}
	return viewModels
}

func StatusViewModelFromStatus(s db.StatusesRow) view.StatusViewModel {
	var id string
	if s.Status.ID != uuid.Nil {
		id = s.Status.ID.String()
	}

	return view.StatusViewModel{
		ID:   id,
		Name: s.Status.Name,
	}
}
