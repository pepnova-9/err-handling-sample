package repository

import (
	"context"
	"errors"
	"fmt"

	domain2 "github.com/pepnova-9/err-handling-sample/sample1/domain"
)

func GetUser(ctx context.Context, userID string) (domain2.User, error) {

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
			return domain2.User{}, domain2.ErrRecordNotFound
		default:
			return domain2.User{}, fmt.Errorf("failed to get user from db: %w", err)
		}
	}

	user := domain2.User{
		ID: userID,
	}

	return user, nil
}
