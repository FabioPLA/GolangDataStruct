package main

import (
	"sort"
)

//算法思想；将num2放到num1的尾部，然后将num1数组排序
func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}
func main() {
	nums1, m, nums2 := []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}
	n := 2
	merge(nums1, m, nums2, n)
}
