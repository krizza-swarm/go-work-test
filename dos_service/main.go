package main

import (
	"fmt"

	"github.com/krizza-swarm/go-work-test/math_lib/operations"
)

func main() {
	fmt.Println("dos_service using Add from module math_lib and package operations")
	fmt.Println(operations.Add(1, 2))
}
