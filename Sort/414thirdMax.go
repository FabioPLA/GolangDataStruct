package main

import (
	"fmt"
	"sort"
)

func thirdMax(nums []int) int {
	sort.Ints(nums)
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}
	res := nums[:j]
	n := len(res)
	if n <= 2 {
		return res[n-1]
	} else {
		return res[n-3]
	}
}
func main() {
	nums := []int{3, 2, 1, 5}
	fmt.Println(thirdMax(nums))
}
