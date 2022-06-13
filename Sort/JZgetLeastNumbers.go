package main

import (
	"fmt"
	"sort"
)

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	result := []int{}
	for i := 0; i < k; i++ {
		result = append(result, arr[i])
	}
	return result
}
func main() {
	arr := []int{0, 1, 2, 1}
	k := 1
	fmt.Println(getLeastNumbers(arr, k))
}
