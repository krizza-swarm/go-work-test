package main

import (
	"fmt"

	"github.com/krizza-swarm/go-work-test/math_lib/display"
	"github.com/krizza-swarm/go-work-test/math_lib/operations"
)

func main() {
	fmt.Println("uno_service")
	sum := operations.Add(1, 2, 3)
	fmt.Println(display.PrintAddition([]int{1, 2, 3}, sum))
}
