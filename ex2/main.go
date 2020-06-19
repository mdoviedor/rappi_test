package main

import "fmt"

type ComplexityUCInterface interface {
	calculate(tag int, iteration int) []int
}

func Handler(complexityUC ComplexityUCInterface, tag int, iteration int) []int {
	return complexityUC.calculate(tag, iteration)
}

func main() {
	cfg := NewConfig()
	sl := NewSl(cfg)

	fmt.Print("input tag: ")
	var tag int
	fmt.Scanf("%d", &tag)

	fmt.Print("input iteration: ")
	var iteration int
	fmt.Scanf("%d", &iteration)

	fmt.Println(Handler(sl.complexityUC, tag, iteration))
}
