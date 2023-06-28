package orm

import "errors"

// ORMが定義しているerrorとしてここで書いてます
var (
	ErrRecordNotFound = errors.New("record not found")
	ErrUnknownError   = errors.New("unknown error")
)
