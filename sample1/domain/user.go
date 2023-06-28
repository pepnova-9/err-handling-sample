package domain

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
)

var (
	ErrInvalidUserNameLength    = errors.New("user name must be shorter than 5 characters")
	ErrInvalidUserNameCharacter = errors.New("user name must contain only alphabet characters")
)

// Domain層でエラーを投げてもControllerでどのリクエストボディのフィールドがエラーになったのか(この場合だと"user_name"とか)をマッピングして返したい場合はどうするか？
// Domain層で入れるのは違う。 ドメインの関心ごとではないものが入るから。あるAPIではuser_nameだけど別n APIではfollower_user_name のようにフィールドが変わる。レスポンスのエラーメッセージもcontrollerで加える必要がありそう。
// 結局このerrorをControllerで受け取って処理する必要がある。
// controllerで共通のerrorハンドリング処理を書く場合、リクエストボディのフィールド名をレスポンスに含めたい時やエラーメッセージを変えたい時に共通化できなくなってくる。
// 一方で各controllerで同じようなエラーハンドリングが出てきてしまう。例えば"ユーザ名は4文字以下である必要があります"みたいな。
// domainのエラーはcontrollerで参照して良いか？愚直にusecaseで別のエラーに変えて渡すべきか？

type User struct {
	ID   string
	Name string
}

func NewUser(name string) (User, error) {
	if len(name) > 4 {
		return User{}, ErrInvalidUserNameLength
	}

	if !isAlphanumericDashUnderscore(name) {
		return User{}, ErrInvalidUserNameCharacter
	}

	return User{
		ID:   uuid.Must(uuid.NewUUID()).String(),
		Name: name,
	}, nil
}

var alphabetRegexp = regexp.MustCompile(`^[a-zA-Z]+$`)

func isAlphanumericDashUnderscore(s string) bool {
	return alphabetRegexp.MatchString(s)
}
