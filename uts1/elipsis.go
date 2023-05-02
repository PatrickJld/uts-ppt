package main

import (
	"fmt"
)

func elipsis(num ...int) int {

	total := 0

	for _, i := range num {
		total += i
	}

	return total
}

func main() {
	fmt.Println(elipsis(2, 4, 6, 8))
	fmt.Println(elipsis(2, 4, 6, 8, 10, 12, 14, 16))
}
