package main

import "fmt"

func swap(a, b int) []int {
	nums := []int{}
	a = a ^ b
	b = a ^ b
	a = a ^ b
	nums = append(nums, a)
	nums = append(nums, b)
	return nums
}
func main() {
	a, b := 10, 3
	fmt.Println(swap(a, b))
}
