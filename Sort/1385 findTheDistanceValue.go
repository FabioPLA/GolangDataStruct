package main

import (
	"math"
	"sort"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)

	count := 0

	for _, i2 := range arr1 {
		r := sort.Search(len(arr2)-1, func(i int) bool {
			switch {
			case i2 > arr2[i]:
				return false
			default:
				return true
			}
		})

		dis := math.MaxInt32
		if r > 0 {
			dis = min(dis, getDis(i2, arr2[r-1]))
		}
		if r < len(arr2)-1 {
			dis = min(dis, getDis(i2, arr2[r+1]))
		}
		dis = min(dis, getDis(i2, arr2[r]))

		if dis <= d {

		} else {
			count++
		}

	}
	return count
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func getDis(x int, y int) int {
	dis := x - y

	if dis < 0 {
		dis = -dis
	}
	return dis
}
func main() {

}
