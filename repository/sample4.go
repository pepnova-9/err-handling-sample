package repository

import (
	"context"
	"fmt"

	"github.com/pepnova-9/err-handling-sample/errwrapper"

	"github.com/pepnova-9/err-handling-sample/domain"
	"github.com/pepnova-9/err-handling-sample/pkg/orm"
)

type Sample4 struct{}

func (s *Sample4) GetUser(ctx context.Context, userID string) (_ domain.User, err error) {
	//defer errwrapper.Wrap(&err, "Sample4.GetUser(%q)", userID)
	defer errwrapper.WrapStack(&err, "Sample4.GetUser(%q)", userID)

	// ===== SQL作ってDBからUserを取得する処理
	// user, err := db...
	// わざとエラー出すためにここでuserIDの値によってエラー変えてます
	switch userID {
	case "not_found":
		err = orm.ErrRecordNotFound
	case "unknown":
		err = orm.ErrUnknownError
	case "repository_panic":
		panic("repository panic")
	default:
		err = nil
	}
	// =====

	if err != nil {
		switch err {
		case orm.ErrRecordNotFound:
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
