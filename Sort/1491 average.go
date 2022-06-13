package main

import (
	"fmt"
)

func average(salary []int) float64 {
	min, max := 1000000, 1000
	var sum int64 = 0
	for _, s := range salary {
		if s > max {
			max = s
		}

		if s < min {
			min = s
		}

		sum += int64(s)
	}

	return float64(sum-int64(min)-int64(max)) / float64(len(salary)-2)
}
func main() {
	salary := []int{6000, 5000, 4000, 3000, 2000, 1000}
	fmt.Println(average(salary))
}
