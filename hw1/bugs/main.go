package main

import (
	"fmt"
)

// вам необходимо поправить баги в этих функциях чтобы тесты в main_test.go проходили

func ExistCounter() int {
	ids := []int{1, 2, 3, 4, 5, 6}
	enabled := map[int]bool{
		1: false,
		2: true,
		4: false,
		6: true,
	}
	totalExist := 0
	for _, id := range ids {
		if _, exist := enabled[id]; exist {
			totalExist++
		}
	}
	return totalExist
}

func Shadowing() int {
	x := 1
	for i := 0; i <= 10; i++ {
		x = x + 1
		x *= 2
	}
	return x
}

func BadMap() (resultErr error) {
	defer func() {
		if err := recover(); err != nil {
			resultErr = fmt.Errorf("recover: %v", err)
		}
	}()

	users := make(map[int]string, 100500)
	users[100500] = "rvasily"

	return nil
}
