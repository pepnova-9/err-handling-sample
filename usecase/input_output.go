package usecase

type UpdateUserInput struct {
	UserID string
	Name   string
}

type UpdateUserOutput struct {
	ID   string
	Name string
}
