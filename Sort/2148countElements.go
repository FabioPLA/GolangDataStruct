package main

import "fmt"

func countElements(nums []int) int {
	mx, mn, mxc, mnc := -100001, 100001, 0, 0
	for _, num := range nums {
		if num < mn {
			mn, mnc = num, 1
		} else if num == mn {
			mnc++
		}
		if num > mx {
			mx, mxc = num, 1
		} else if num == mx {
			mxc++
		}
	}
	if mn == mx {
		return 0
	}
	return len(nums) - mxc - mnc
}
func main() {
	nums := []int{11, 7, 2, 15}
	fmt.Println(countElements(nums))
}
