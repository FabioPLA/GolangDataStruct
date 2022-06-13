package main

import (
	"fmt"
)

func dominantIndex(nums []int) int {
	m1, m2, idx := -1, -1, 0
	for i, num := range nums {
		if num > m1 {
			m1, m2, idx = num, m1, i
		} else if num > m2 {
			m2 = num
		}
	}
	if m1 >= m2*2 {
		return idx
	}
	return -1
}
func main() {
	nums := []int{3, 6, 1, 0}
	fmt.Println(dominantIndex(nums))
}
