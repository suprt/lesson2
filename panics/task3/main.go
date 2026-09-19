package main

import "fmt"

func Level1() {
	defer func() {
		if pan := recover(); pan != nil {
			fmt.Printf("Паника обработана на уровне 1: %v \n", pan)
		}
	}()
	Level2()
}

func Level2() {
	defer fmt.Println("Завершаем Level2")
	Level3()
}

func Level3() {
	panic("ошибка в Level3")
}

func main() {
	Level1()
}
