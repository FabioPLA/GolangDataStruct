package main

import "fmt"

func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}
	return nums
}
func main() {
	nums := []int{10, 5, 11, 9}
	fmt.Println(bubbleSort(nums))
}
