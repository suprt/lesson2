package main

import "fmt"

func handle() error {
	return fmt.Errorf("this is an error")
}

func main() {
	fmt.Println(handle())
}
