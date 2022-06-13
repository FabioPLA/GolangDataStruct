package main

import (
	"fmt"
	"sort"
)

func minimumCost(cost []int) int {
	ans := 0
	sort.Ints(cost)
	for i := len(cost) - 1; i >= 0; i -= 3 {
		ans += cost[i]
		if i > 0 {
			ans += cost[i-1]
		}
	}
	return ans
}
func main() {
	cost := []int{6, 5, 7, 9, 2, 2}
	fmt.Println(minimumCost(cost))
}
