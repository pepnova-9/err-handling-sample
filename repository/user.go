package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/pepnova-9/err-handling-sample/domain"
)

func GetUser(ctx context.Context, userID string) (domain.User, error) {

	// SQL作ってDBからUserを取得する処理
	// user, err := db...

	// わざとエラー出すためにここでuserIDの値によってエラー変えてます
	var err error
	switch userID {
	case "not_found":
		err = ORMErrRecordNotFound
	case "unknown":
		err = errors.New("unknown error")
	default:
		err = nil
	}

	if err != nil {
		switch err {
		case ORMErrRecordNotFound:
			return domain.User{}, domain.ErrRecordNotFound
		default:
			return domain.User{}, fmt.Errorf("failed to get user from db: %w", err)
		}
	}

	user := domain.User{
		ID: userID,
	}

	return user, nil
}
