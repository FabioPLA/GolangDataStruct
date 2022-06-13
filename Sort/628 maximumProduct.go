package main

import (
	"fmt"
	"sort"
)

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return max(nums[0]*nums[1]*nums[n-1], nums[n-3]*nums[n-2]*nums[n-1])
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(maximumProduct(nums))
}
