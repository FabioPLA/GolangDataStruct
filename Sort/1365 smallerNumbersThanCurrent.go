package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {

	ans := []int{}
	for _, k := range nums {
		cnt := 0
		for _, v := range nums {
			if v < k {
				cnt++
			}
		}
		ans = append(ans, cnt)
	}
	return ans
}
func main() {
	nums := []int{8, 1, 2, 2, 3}
	fmt.Println(smallerNumbersThanCurrent(nums))
}
