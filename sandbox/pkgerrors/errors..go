package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	err := Error1()
	fmt.Printf("%+v\n", err)
}

func Error1() error {
	err := Error2()
	err = errors.Wrap(err, "This is Error1")
	return err
}

func Error2() error {
	err := RootError()
	err = errors.Wrap(err, "This is Error2")
	return err
}

func RootError() error {
	err := errors.New("This is RootError")
	return err
}
