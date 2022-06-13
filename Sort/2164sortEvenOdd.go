package main

import "fmt"

func sortEvenOdd(nums []int) []int {
	for i := 2; i < len(nums); i = i + 2 {
		for j := 0; j < i; j = j + 2 {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	for i := 3; i < len(nums); i = i + 2 {
		for j := 1; j < i; j = j + 2 {
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}
func main() {
	nums := []int{4, 1, 2, 3}
	fmt.Println(sortEvenOdd(nums))
}
