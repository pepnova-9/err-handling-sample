package main

import (
	"fmt"

	"github.com/cockroachdb/errors"
)

func ErrorExp() error {
	err := errors.New("some error")
	return err
}

func main() {
	err := ErrorExp()
	fmt.Printf("%+v\n", err)
}
