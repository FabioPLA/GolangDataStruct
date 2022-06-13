//输入：arr = [40,10,20,30]
//输出：[4,1,2,3]
//算法：哈希表+两遍遍历 ，哈希表记录位置
package main

import (
	"fmt"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	arrclone := make([]int, len(arr))
	copy(arrclone, arr)
	sort.Ints(arr)
	var result []int
	seq := 1
	resultMap := make(map[int]int)
	for _, v := range arr {
		if _, ok := resultMap[v]; ok {
			continue
		}
		resultMap[v] = seq
		seq++
		fmt.Println(resultMap)
	}
	for _, v := range arrclone {
		result = append(result, resultMap[v])
	}
	return result
}
func main() {
	arr := []int{40, 10, 20, 30}
	fmt.Println(arr)
	fmt.Println(arrayRankTransform(arr))
}
