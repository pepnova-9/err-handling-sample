package orm

import "errors"

// ORMが定義しているerrorとしてここで書いてます
var (
	ErrRecordNotFound = errors.New("db record not found error")
	ErrUnknownError   = errors.New("db unknown error")
)
