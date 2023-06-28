package usecase

import (
	"context"
	"fmt"

	"github.com/pepnova-9/err-handling-sample/domain"
	"github.com/pepnova-9/err-handling-sample/repository"
)

type Sample1 struct{}

type GetUserOutput struct {
	ID   string
	Name string
}

type UpdateUserInput struct {
	UserID string
	Name   string
}

type UpdateUserOutput struct {
	ID   string
	Name string
}

func (s *Sample1) UpdateUserUsecase(ctx context.Context, input UpdateUserInput) (UpdateUserOutput, error) {
	sample1 := repository.Sample1{}

	// 権限エラーの場合
	if input.UserID == "unauthorized" {
		return UpdateUserOutput{}, ErrUnauthorized
	}

	user, err := sample1.GetUser(ctx, input.UserID)
	if err != nil {
		switch err {
		case domain.ErrRecordNotFound:
			return UpdateUserOutput{}, ErrUserNotFound
		default:
			return UpdateUserOutput{}, fmt.Errorf("failed to repository.GetUser: %w, %w", ErrUnexpectedError, err)
		}
	}

	err = user.ChangeName(input.Name)
	if err != nil {
		return UpdateUserOutput{}, fmt.Errorf("failed to User.ChangeName: %w", err)
	}

	return UpdateUserOutput{ID: user.ID, Name: user.Name}, nil
}
