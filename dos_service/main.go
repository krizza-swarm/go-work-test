package main

import (
	"fmt"

	"github.com/krizza-swarm/go-work-test/math_lib/display"
	"github.com/krizza-swarm/go-work-test/math_lib/operations"
)

func main() {
	fmt.Println("dos_service")
	sum := operations.Add(4, 5, 6)
	fmt.Println(display.PrintAddition([]int{4, 5, 6}, sum))
}
