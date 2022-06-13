package main

import (
	"fmt"
	"sort"
)

func findFinalValue(nums []int, original int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if original == nums[i] {
			original *= 2
		}
	}
	return original
}
func main() {
	nums := []int{5, 3, 6, 1, 12}
	original := 3
	fmt.Println(findFinalValue(nums, original))
}
