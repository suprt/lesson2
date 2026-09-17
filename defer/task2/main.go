package main

import "fmt"

func main() {
	value := 123
	defer func() {
		fmt.Println(value)
	}()
	changeValue(&value)
}
func changeValue(value *int) {
	*value = 456
}

/*
вывод до исправления - 123
defer захватывает значение аргументов в момент объявления
исправление -> defer анонимной функции, содержание функции defer не анализируется
*/
