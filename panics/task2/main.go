package main

import "fmt"

// SafeDivide
//  При панике будет использован zero value возвращаемого типа, то есть 0. Можно было бы
//  использовать именованный возврат.
func SafeDivide(a, b int) int {
	defer func() {
		recover()
	}()
	if b == 0 {
		panic("деление на ноль")
	}
	return a / b
}

func main() {
	fmt.Println(SafeDivide(10, 2)) // Ожидаемый результат: 5
	fmt.Println(SafeDivide(10, 0)) // Ожидаемый результат: 0 (без паники)
}
