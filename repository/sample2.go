package repository

import (
	"context"

	"github.com/pkg/errors"

	"github.com/pepnova-9/err-handling-sample/domain"
	"github.com/pepnova-9/err-handling-sample/pkg/orm"
)

type Sample2 struct{}

func (s *Sample2) GetUser(ctx context.Context, userID string) (domain.User, error) {
	// ===== SQL作ってDBからUserを取得する処理
	// user, err := db...
	// わざとエラー出すためにここでuserIDの値によってエラー変えてます
	var err error
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
			return domain.User{}, errors.Wrapf(domain.ErrRecordNotFound, "no record found userID=%s", userID)
		default:
			return domain.User{}, errors.Wrap(err, "failed to get user from db:")
		}
	}

	user := domain.User{
		ID: userID,
	}

	return user, nil
}
