package usecase

import (
	"context"
	"errors"

	pkgErrors "github.com/pkg/errors"

	"github.com/pepnova-9/err-handling-sample/domain"
	"github.com/pepnova-9/err-handling-sample/repository"
)

type Sample2 struct{}

func (s *Sample2) UpdateUserUsecase(ctx context.Context, input UpdateUserInput) (UpdateUserOutput, error) {
	sampleRepo := repository.Sample2{}

	// 権限エラーの場合
	if input.UserID == "unauthorized" {
		return UpdateUserOutput{}, pkgErrors.Wrapf(ErrUnauthorized, "userID=%s:", input.UserID)
	}

	user, err := sampleRepo.GetUser(ctx, input.UserID)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrRecordNotFound):
			// "failed to repository.GetUser:" と入れるか・・・？けど"DB"のエラーではなく、レコードがないことをerrorとしてハンドリングしているだけなので違和感ある。"not found"とかだと重複するし。空文字もありか？
			return UpdateUserOutput{}, pkgErrors.Wrap(ErrUserNotFound, "")
		default:
			// Joinしちゃうと+v出力の時に下位のerrorのスタックトレースが取れなくなるので注意。 Joinの第一引数をerrors.withStackにしないとダメかも。
			//return UpdateUserOutput{}, pkgErrors.Wrap(errors.Join(err, ErrUnexpectedError), "failed to repository.GetUser:")
			return UpdateUserOutput{}, pkgErrors.Wrap(err, "failed to repository.GetUser:")
		}
	}

	err = user.ChangeName(input.Name)
	if err != nil {
		// ここでラップしてもドメイン層ではpkg/errorsでラップしてないのでドメインレイヤーのどの行でエラーになったのかはスタックトレースからはわからない。
		// ドメインレイヤーでもpkg/errorsでラップするか・・・？
		return UpdateUserOutput{}, pkgErrors.Wrap(err, "failed to User.ChangeName:")
	}

	// DBへ更新

	return UpdateUserOutput{ID: user.ID, Name: user.Name}, nil
}
