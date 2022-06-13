package main

import (
	"fmt"
	"sort"
)

func exchange(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i]%2 > nums[j]%2
	})
	return nums
}
func main() {
	nums := []int{}
	fmt.Println(exchange(nums))
}
