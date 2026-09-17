package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	for i := 1; i < 4; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("end")
}

/*
end
4
3
2
1
*/
