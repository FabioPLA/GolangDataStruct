package main

import (
	"fmt"
	"sort"
)

//算法思想：将数组排序之后，即可根据数组中每个下标处的元素是否和下标相等，得到丢失的数字。
func missingnumber(nums []int) int {
	sort.Ints(nums)
	for i, num := range nums {
		if num != i {
			return i
		}
	}
	return len(nums)
}
func main() {
	nums := []int{3, 0, 1}
	fmt.Println(missingnumber(nums))
}
