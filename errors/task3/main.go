package main

import (
	"errors"
	"fmt"
	"math/rand"
)

var (
	ErrNotFound  = errors.New("ресурс не найден")
	TimeoutError = errors.New("таймаут операции")
)

func SimulateRequest() error {
	r := rand.Intn(100)
	if r < 50 {
		return fmt.Errorf("запрос не выполнен: %w", TimeoutError)
	}
	if r < 80 {
		return fmt.Errorf("ошибка: %w", ErrNotFound)
	}
	return errors.New("неизвестная ошибка")
}

func ProcessError(err error) {
	if errors.Is(err, TimeoutError) {
		fmt.Println("Требуется повторная попытка")
	} else if errors.Is(err, ErrNotFound) {
		fmt.Println("Ресурс не найден")
	} else {
		fmt.Println("Неизвестная ошибка")
	}

}

func main() {
	ProcessError(SimulateRequest())
}
