package repository

import "errors"

var (
	ORMErrRecordNotFound = errors.New("record not found") // ORMが定義しているerrorとしてここで書いてます
)
