package main

import "sort"

func specialArray(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	l, r := 0, n
	for l < r {
		// 加 1 避免陷入死循环
		m := (l + r + 1) >> 1
		// 大于等于可能是特征值
		if nums[n-m] >= m {
			l = m
			//	但小于一定不是特征值
		} else {
			r = m - 1
		}
	}
	// 应该有 l 个数字大于等于 l 则 nums[n-l] >= l
	// 且对于任意 x < n-l 应有 nums[x] < l
	idx := n - l
	if idx > 0 && nums[idx-1] >= l {
		return -1
	}
	return l
}
func main() {

}
