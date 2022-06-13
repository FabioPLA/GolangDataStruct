package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	sorted := append([]int{}, heights...)
	sort.Ints(heights)
	count := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != sorted[i] {
			count++
		}
	}
	return count
}
func main() {
	heights := []int{5, 1, 2, 3, 4}
	fmt.Println(heightChecker(heights))
}
