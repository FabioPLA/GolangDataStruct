package main

import "fmt"

func sortArrayByParity(nums []int) []int {
	result := []int{}
	for _, num := range nums {
		if num%2 == 0 {
			result = append(result, num)
		}
	}
	for _, num := range nums {
		if num%2 == 1 {
			result = append(result, num)
		}
	}
	return result
}
func main() {
	nums := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(nums))
}
