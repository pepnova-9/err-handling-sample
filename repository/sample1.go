package repository

import (
	"context"
	"fmt"

	"github.com/pepnova-9/err-handling-sample/pkg/orm"

	"github.com/pepnova-9/err-handling-sample/domain"
)

type Sample1 struct{}

func (s *Sample1) GetUser(ctx context.Context, userID string) (domain.User, error) {

	// ===== SQL作ってDBからUserを取得する処理
	// user, err := db...
	// わざとエラー出すためにここでuserIDの値によってエラー変えてます
	var err error
	switch userID {
	case "not_found":
		err = orm.ErrRecordNotFound
	case "unknown":
		err = orm.ErrUnknownError
	default:
		err = nil
	}
	// =====

	if err != nil {
		switch err {
		case orm.ErrRecordNotFound:
			return domain.User{}, domain.ErrRecordNotFound
		default:
			// 外部ライブラリのerrorは内部エラーでラップする
			return domain.User{}, fmt.Errorf("failed to get user from db: %w", err)
		}
	}

	user := domain.User{
		ID: userID,
	}

	return user, nil
}
