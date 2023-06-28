package usecase

import (
	"context"
	"fmt"

	"github.com/pepnova-9/err-handling-sample/sample1/repository"

	domain2 "github.com/pepnova-9/err-handling-sample/sample1/domain"
)

type GetUserOutput struct {
	ID   string
	Name string
}

func GetUserUsecase(ctx context.Context, userID string) (GetUserOutput, error) {
	// 権限エラーの場合
	if userID == "unauthorized" {
		return GetUserOutput{}, ErrUnauthorized
	}

	user, err := repository.GetUser(ctx, userID)
	if err != nil {
		switch err {
		case domain2.ErrRecordNotFound:
			return GetUserOutput{}, ErrUserNotFound
		default:
			return GetUserOutput{}, fmt.Errorf("failed to repository.GetUser: %w, %w", ErrUnexpectedError, err)
		}
	}

	return GetUserOutput{
		ID: user.ID,
	}, nil
}

type CreateUserInput struct {
	Name string
}

type CreateUserOutput struct {
	ID   string
	Name string
}

func CreateUserUsecase(ctx context.Context, input CreateUserInput) (CreateUserOutput, error) {

	user, err := domain2.NewUser(input.Name)
	if err != nil {
		return CreateUserOutput{}, fmt.Errorf("domain.NewUser: %w", err)
	}

	return CreateUserOutput{ID: user.ID, Name: user.Name}, nil
}
