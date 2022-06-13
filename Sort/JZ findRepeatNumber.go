package main

import (
	"fmt"
	"sort"
)

func findRepeatNumber(nums []int) int {
	result := 0
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			result = nums[i]
		}
	}
	return result
}
func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(nums))
}
