//输入：arr = [4,2,1,3]
//输出：[[1,2],[2,3],[3,4]]
package main

import (
	"fmt"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	length := len(arr)
	var result [][]int
	min := arr[length-1] - arr[0]
	for i := 1; i < length; i++ {
		cha := arr[i] - arr[i-1]
		if min > cha {
			result = [][]int{}
			min = cha
			result = append(result, []int{arr[i-1], arr[i]})
		} else if min == cha {
			result = append(result, []int{arr[i-1], arr[i]})
		}
	}
	return result
}
func main() {
	arr := []int{4, 2, 1, 3}
	fmt.Println(minimumAbsDifference(arr))
}
