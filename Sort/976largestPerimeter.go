package main

import (
	"fmt"
	"sort"
)

func largestPerimeter(a []int) int {
	sort.Ints(a)
	for i := len(a) - 1; i >= 2; i-- {
		if a[i-2]+a[i-1] > a[i] {
			return a[i-2] + a[i-1] + a[i]
		}
	}
	return 0
}
func main() {
	nums := []int{2, 1, 2}
	fmt.Println(largestPerimeter(nums))
}
