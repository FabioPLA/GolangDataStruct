package main

import (
	"fmt"
	"sort"
)

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	for i := 1; i < len(arr)-1; i++ {
		if arr[i]*2 != arr[i-1]+arr[i+1] {
			return false
		}
	}
	return true
}

func main() {
	arr := []int{3, 5, 1}
	fmt.Println(canMakeArithmeticProgression(arr))
}
