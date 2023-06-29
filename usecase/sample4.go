package usecase

import (
	"context"
	"errors"

	"github.com/pepnova-9/err-handling-sample/errwrapper"

	"github.com/pepnova-9/err-handling-sample/domain"
	"github.com/pepnova-9/err-handling-sample/repository"
)

type Sample4 struct{}

func (s *Sample4) UpdateUserUsecase(ctx context.Context, input UpdateUserInput) (_ UpdateUserOutput, err error) {
	defer errwrapper.Wrap(&err, "Sample4.UpdateUserUsecase(%#v)", input)

	sampleRepo := repository.Sample4{}

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
			return UpdateUserOutput{}, err
		}
	}

	err = user.ChangeName(input.Name)
	if err != nil {
		return UpdateUserOutput{}, err
	}

	// DBへ更新

	return UpdateUserOutput{ID: user.ID, Name: user.Name}, nil
}
