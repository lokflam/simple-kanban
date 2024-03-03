package kanban

import (
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lokflam/simple-kanban/internal/kanban/db"
	"github.com/lokflam/simple-kanban/internal/kanban/view"
)

func Router(p *pgxpool.Pool) http.Handler {
	q := db.New(p)

	r := chi.NewRouter()
	r.Use(middleware.SetHeader("Content-Type", "text/html; charset=utf-8"))
	r.Get("/", BoardHandler(q))
	r.Put("/cards", UpsertCardHandler(q))
	r.Delete("/cards/{id}", DeleteCardHandler(q))
	r.Get("/card-form", CardFormHandler(q))
	r.Get("/card-form/{id}", CardFormHandler(q))

	return r
}

func BoardHandler(q db.Querier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hxRequest := r.Header.Get("Hx-Request")
		hxTarget := r.Header.Get("Hx-Target")

		if hxRequest != "true" {
			component := view.Page(view.Board(nil, true))
			if err := component.Render(r.Context(), w); err != nil {
				RenderError(w, (&ErrRenderFailed{err}))
				return
			}
			return
		}

		statuses, err := q.Statuses(r.Context())
		if err != nil {
			RenderError(w, (&ErrQueryFailed{err}))
			return
		}

		boardVM := view.BoardViewModel{}
		for _, s := range statuses {
			statusVM := view.StatusViewModel{
				ID:   s.Status.ID.String(),
				Name: s.Status.Name,
			}

			boardVM[statusVM] = []view.CardViewModel{}

			cards, err := q.CardsByStatus(r.Context(), s.Status.ID)
			if err != nil {
				RenderError(w, (&ErrQueryFailed{err}))
				return
			}

			for _, c := range cards {
				boardVM[statusVM] = append(boardVM[statusVM], CardViewModelFromCard(c))
			}
		}

		board := view.Board(boardVM, false)

		var component templ.Component

		switch hxTarget {
		case "board":
			component = board
		default:
			component = view.Page(board)
		}

		if err := component.Render(r.Context(), w); err != nil {
			RenderError(w, (&ErrRenderFailed{err}))
			return
		}
	}
}

func CardFormHandler(q db.Querier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")

		var card db.Card
		if idParam != "" {
			id, err := uuid.Parse(idParam)
			if err != nil {
				RenderError(w, (&ErrInvalidID{err}))
				return
			}

			card, err = q.Card(r.Context(), id)
			if err != nil {
				RenderError(w, (&ErrQueryFailed{err}))
				return
			}
		}

		statuses, err := q.Statuses(r.Context())
		if err != nil {
			RenderError(w, (&ErrQueryFailed{err}))
			return
		}

		component := view.CardFormDialog(view.CardFormViewModel{
			Open:     true,
			Card:     CardViewModelFromCard(card),
			Statuses: StatusViewModelsFromListStatusesRows(statuses),
		})
		if err := component.Render(r.Context(), w); err != nil {
			RenderError(w, (&ErrRenderFailed{err}))
			return
		}
	}
}

func UpsertCardHandler(q db.Querier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dryRun := r.URL.Query().Get("dry_run") == "true"
		hxTarget := r.Header.Get("Hx-Target")

		cardFields := &CardForm{}
		if err := cardFields.Bind(r); err != nil {
			RenderError(w, (&ErrInvalidRequestData{err}))
			return
		}

		vErr := cardFields.Validate()
		if vErr != nil || dryRun {
			statuses, err := q.Statuses(r.Context())
			if err != nil {
				RenderError(w, (&ErrQueryFailed{err}))
				return
			}

			vm := view.CardFormViewModel{
				Open: true,
				Card: view.CardViewModel{
					ID:       cardFields.ID.String(),
					StatusID: cardFields.StatusID.String(),
					Title:    cardFields.Title,
					Content:  cardFields.Content,
				},
				Statuses: StatusViewModelsFromListStatusesRows(statuses),
			}

			if vErr != nil {
				vm.FieldErrors = vErr.(*ErrInvalidFields).FieldErrors
			}

			var component templ.Component

			switch hxTarget {
			case "card-title-field":
				if _, ok := vm.FieldErrors["Title"]; ok {
					w.WriteHeader(http.StatusBadRequest)
				}
				component = view.CardFormTitleField(vm)

			case "card-content-field":
				if _, ok := vm.FieldErrors["Content"]; ok {
					w.WriteHeader(http.StatusBadRequest)
				}
				component = view.CardFormContentField(vm)

			default:
				if len(vm.FieldErrors) > 0 {
					w.WriteHeader(http.StatusBadRequest)
				}
				component = view.CardFormDialog(vm)
			}

			if err := component.Render(r.Context(), w); err != nil {
				RenderError(w, (&ErrRenderFailed{err}))
				return
			}
			return
		}

		id := cardFields.ID
		if cardFields.ID == uuid.Nil {
			newID, err := uuid.NewV7()
			if err != nil {
				RenderError(w, (&ErrGenerateIDFailed{err}))
				return
			}
			id = newID
		}

		ts := time.Now()

		err := q.UpsertCard(r.Context(), db.UpsertCardParams{
			ID:        id,
			Title:     cardFields.Title,
			Content:   cardFields.Content,
			StatusID:  cardFields.StatusID,
			CreatedAt: ts,
			UpdatedAt: ts,
		})
		if err != nil {
			RenderError(w, (&ErrQueryFailed{err}))
			return
		}

		w.Header().Add("Hx-Reswap", "none")
		w.Header().Add("Hx-Trigger", "card-update")
		if cardFields.ID == uuid.Nil {
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

func DeleteCardHandler(q db.Querier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			RenderError(w, (&ErrInvalidID{err}))
			return
		}

		err = q.DeleteCard(r.Context(), id)
		if err != nil {
			RenderError(w, (&ErrQueryFailed{err}))
			return
		}

		w.Header().Add("Hx-Reswap", "none")
		w.Header().Add("Hx-Trigger", "card-update")
		w.WriteHeader(http.StatusOK)
	}
}
