package view

type BoardViewModel map[StatusViewModel][]CardViewModel

type StatusViewModel struct {
	ID   string
	Name string
}

type CardViewModel struct {
	ID       string
	Title    string
	Content  string
	StatusID string
}

type CardFormViewModel struct {
	Open        bool
	Card        CardViewModel
	FieldErrors map[string][]error
	Statuses    []StatusViewModel
}
