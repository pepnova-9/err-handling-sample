package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Code     int
	Message  string
	childErr error
}

func NewMyError(code int, msg string, childErr error) *MyError {
	return &MyError{
		Code:     code,
		Message:  msg,
		childErr: childErr,
	}
}

var Err101Error = NewMyError(101, "my error", nil)

func raise101Error() error {
	return Err101Error
}

func (e *MyError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

//func (e *MyError) Is(target error) bool {
//	if t, ok := target.(*MyError); ok {
//		return e.Code == t.Code
//	}
//	return false
//}

//func (e *MyError) Unwrap() error { // ★ errors.Unwrap()に関係するメソッド（後述）
//	return e.childErr
//}

func ErrorHandling1() {
	fmt.Println("=== ErrorHandling1 ===")
	err := raise101Error()
	fmt.Printf("%v \n", err)

	// errors.Isでエラーの型を判定できる
	// これでは拾えない
	if errors.Is(err, &MyError{}) {
		fmt.Printf("&MyErrorで比較: 同じエラー \n")
	} else {
		fmt.Printf("&MyErrorで比較: 違うエラー \n")
	}

	// errors.Isでエラーの型を判定できる
	// これで拾える
	if errors.Is(err, &MyError{Code: 101}) {
		fmt.Printf("&MyError{Code: 101}で比較: 同じエラー \n")
	} else {
		fmt.Printf("&MyError{Code: 101}で比較: 違うエラー \n")
	}

	// 宣言したエラー型で拾える
	if errors.Is(err, Err101Error) {
		fmt.Printf("Err101Errorです \n")
	} else {
		fmt.Printf("Err101Errorじゃないです。 \n")
	}

	// MyErrorのfieldを使いたい時はerrors.Asでキャストする
	myError := &MyError{}
	if ok := errors.As(err, &myError); ok {
		fmt.Printf("Code: %v \n", myError.Code)
		fmt.Printf("Message: %v \n", myError.Message)
	}
}

var OtherPackageErr = errors.New("other package error")

func ErrorHandling2() {
	fmt.Println("=== ErrorHandling2 ===")
	myError := NewMyError(202, "my error having child error", OtherPackageErr)

	if errors.Is(myError, OtherPackageErr) {
		fmt.Printf("OtherPackageErrが含まれている \n")
	} else {
		fmt.Printf("OtherPackageErrが含まれていない \n")
	}
}

func main() {
	ErrorHandling1()
	ErrorHandling2()
}
