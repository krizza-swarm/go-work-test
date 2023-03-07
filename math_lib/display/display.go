package display

import (
	"fmt"
	"strconv"
)

func displayOperation(nums []int, symbol string, total int) string {
	opStr := ""
	for k, v := range nums {
		if len(nums) == k+1 {
			opStr += strconv.Itoa(v)
			break
		}
		opStr += fmt.Sprintf("%v %v ", v, symbol)
	}
	return fmt.Sprintf("%v = %v", opStr, total)
}

func PrintAddition(nums []int, total int) string {
	return displayOperation(nums, "+", total)
}

func PrintSubtraction(nums []int, total int) string {
	return displayOperation(nums, "-", total)
}
