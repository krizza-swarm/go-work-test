package operations

import "fmt"

func Add(nums ...int) int {
	fmt.Println("Now using Add function")
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
