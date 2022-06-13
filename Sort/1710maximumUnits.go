package main

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) (sum int) {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	for _, v := range boxTypes {
		size := min(truckSize, v[0])
		sum += size * v[1]
		truckSize -= size
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
