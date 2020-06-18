package main

import (
	"fmt"
)

func Handler(x int, y int) int {
	return x * (2 + y)
}

func main() {
	fmt.Print("input x: ")
	var x int
	fmt.Scanf("%d", &x)

	fmt.Print("input y: ")
	var y int
	fmt.Scanf("%d", &y)

	result := Handler(x, y)
	fmt.Println(fmt.Sprintf("result: %d", result))
}
