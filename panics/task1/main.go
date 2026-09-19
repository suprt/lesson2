package main

import "fmt"

func CausePanic() {
	panic("что-то пошло не так!")
}

func HandlePanic() {
	defer func() {
		if pan := recover(); pan != nil {
			fmt.Printf("Паника перехвачена: %v \n", pan)
		}
	}()

	CausePanic()
}
func main() {
	//CausePanic()
	HandlePanic()
}
