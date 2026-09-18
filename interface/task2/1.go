package main

import "fmt"

type MyError struct {
	data string
}

func (m *MyError) Error() string {
	return m.data
}
func foo(i int) error {
	var err *MyError
	if i > 5 {
		err = &MyError{data: "i>5"}
	}
	return err
}

func main() {
	err := foo(4)
	if err != nil {
		fmt.Println("oops")
	} else {
		fmt.Println("ok")
	}
}

/*
Я думал вывод ok

На самом деле
Вывод - oops

Почему я ошибся - я не учел, что чтобы интерфейс считался nil и тип и значение должно быть nil,
при объявлении var err *MyError мы уже присваиваем тип и интерфейс автоматически перестаёт быть nil

*/
