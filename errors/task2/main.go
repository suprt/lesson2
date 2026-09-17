package main

import (
	"errors"
	"fmt"
)

func SimpleError() error {
	return errors.New("простая ошибка")
}

func FormattedError(age int) error {
	return fmt.Errorf("%w: возраст %d недопустим", errors.New("ошибка"), age)
}

type MyError struct {
	Code int
	Msg  string
}

func (e MyError) Error() string {
	return e.Msg
}
func StructError() error {
	return MyError{Code: 404, Msg: "не найдено"}
}

func main() {
	fmt.Println(SimpleError())
	fmt.Println(FormattedError(0))
	fmt.Println(StructError())
	fmt.Println(StructError().(MyError).Code, StructError().(MyError).Msg)
}
