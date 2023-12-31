package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/pepnova-9/err-handling-sample/domain"
	"github.com/pepnova-9/err-handling-sample/repository"
)

type Sample1 struct{}

func (s *Sample1) UpdateUserUsecase(ctx context.Context, input UpdateUserInput) (UpdateUserOutput, error) {
	sampleRepo := repository.Sample1{}

	// 権限エラーの場合
	if input.UserID == "unauthorized" {
		return UpdateUserOutput{}, ErrUnauthorized
	}

	user, err := sampleRepo.GetUser(ctx, input.UserID)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrRecordNotFound):
			return UpdateUserOutput{}, ErrUserNotFound
		default:
			return UpdateUserOutput{}, fmt.Errorf("failed to repository.GetUser: %w: %w", ErrUnexpectedError, err)
		}
	}

	err = user.ChangeName(input.Name)
	if err != nil {
		return UpdateUserOutput{}, fmt.Errorf("failed to User.ChangeName: %w", err)
	}

	// DBへ更新

	return UpdateUserOutput{ID: user.ID, Name: user.Name}, nil
}
