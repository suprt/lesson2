package main

import (
	"errors"
	"fmt"
)

func main() {
	println("Case 1")
	case1()
	println()
	println()

	println("Case 2")
	case2()
	println()
	println()

	println("Case 3")
	case3()
	println()
	println()

}

func case1() {
	helperWithDefer := func(isError bool) error {
		var retVal error

		defer func() {
			retVal = errors.New("Extra error")
		}()

		if isError {
			retVal = errors.New("Default error")
		}

		return retVal
	}

	helperWithoutDefer := func(isError bool) error {
		var retVal error

		if isError {
			retVal = errors.New("Default error")
		}

		return retVal
	}

	fmt.Println("\twithout:")
	fmt.Println(helperWithoutDefer(false))
	fmt.Println(helperWithoutDefer(true))
	fmt.Println("\twith:")
	fmt.Println(helperWithDefer(false))
	fmt.Println(helperWithDefer(true))
}

/*
вывод -
	without:
<nil>
Default error
	with:
<nil>
Default error

defer выполняется после окончания функции; функция уже вернула значение, а затем выполнился defer

*/

func case2() {
	helperWithDefer := func(isError bool) (retVal error) {
		defer func() {
			retVal = errors.New("Extra error")
		}()

		if isError {
			retVal = errors.New("Default error")
		}

		return
	}

	helperWithoutDefer := func(isError bool) (retVal error) {
		if isError {
			retVal = errors.New("Default error")
		}

		return
	}

	fmt.Println("\twithout:")
	fmt.Println(helperWithoutDefer(false))
	fmt.Println(helperWithoutDefer(true))
	fmt.Println("\twith:")
	fmt.Println(helperWithDefer(false))
	fmt.Println(helperWithDefer(true))
}

/*
Как я думал - поведение аналогично case1
вывод -
	without:
<nil>
Default error
	with:
<nil>
Default error
Как вышло - поведение отличается, вывод
	without:
<nil>
Default error
	with:
Extra error
Extra error
Почему я ошибся - я думал, что defer выполняется после выхода из функции.
На самом деле он выполняется после return, либо же в момент выхода из функции, и может влиять на именованные возвращаемые параметры
*/

func case3() {
	helperWithDefer := func(isError bool) (retVal error) {
		defer func() {
			retVal = errors.New("First Error")
		}()

		defer func() {
			retVal = errors.New("Second Error")
		}()

		if isError {
			retVal = errors.New("Default error")
		}

		return
	}

	helperWithoutDefer := func(isError bool) (retVal error) {
		if isError {
			retVal = errors.New("Default error")
		}

		return
	}

	fmt.Println("\twithout:")
	fmt.Println(helperWithoutDefer(false))
	fmt.Println(helperWithoutDefer(true))
	fmt.Println("\twith:")
	fmt.Println(helperWithDefer(false))
	fmt.Println(helperWithDefer(true))
}

/*
вывод -
	without:
<nil>
Default error
	with:
First Error
First Error

Случай аналогичный case2 - defer изменяет именованный возвращаемый параметр,
defer выполняются в LIFO порядке, поэтому выводится First Error
*/
