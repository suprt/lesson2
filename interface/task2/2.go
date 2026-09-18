package main

import (
	"fmt"
)

type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}

func checkErr(err error) {
	fmt.Println(err == nil)
}

func main() {
	var e1 error
	checkErr(e1)

	var e *errorString
	checkErr(e)

	e = &errorString{}
	checkErr(e)

	e = nil
	checkErr(e)
}

/*
Вывод -
true
false
false
false

e1 переменная относящаяся к интерфейсу error, но не имеющая типа. Соответственно err==nil - true
e имеет тип *errorString, поэтому будет false
e=nil изменяет только значение переменной, но не её тип, поэтому тоже false
*/
