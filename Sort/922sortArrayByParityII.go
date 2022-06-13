package main

import "fmt"

func sortArrayByParityII(nums []int) []int {
	ans := make([]int, len(nums))
	i := 0
	for _, v := range nums {
		if v%2 == 0 {
			ans[i] = v
			i += 2
		}
	}
	i = 1
	for _, v := range nums {
		if v%2 == 1 {
			ans[i] = v
			i += 2
		}
	}
	return ans
}

func main() {
	nums := []int{4, 2, 5, 7}
	fmt.Println(sortArrayByParityII(nums))
}
