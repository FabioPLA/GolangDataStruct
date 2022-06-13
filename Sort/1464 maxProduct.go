package main

import (
	"fmt"
	"sort"
)

func maxProduct(nums []int) int {
	sort.Ints(nums)
	return (nums[len(nums)-1] - 1) * (nums[len(nums)-2] - 1)
}
func main() {
	nums := []int{1, 5, 4, 5}
	fmt.Println(maxProduct(nums))
}
