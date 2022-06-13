package main

import (
	"fmt"
	"sort"
)

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		x, y := arr[i], arr[j]
		//fmt.Println("X", x)
		//fmt.Println("y", y)
		cx, cy := oneCount(x), oneCount(y)
		return cx < cy || cx == cy && x < y
	})
	return arr
}
func oneCount(num int) int {
	ones := 0
	for ; num > 0; num &= (num - 1) {
		ones++
	}
	return ones
}
func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(sortByBits(arr))
}
