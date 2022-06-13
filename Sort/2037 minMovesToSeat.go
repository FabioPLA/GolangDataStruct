package main

import (
	"fmt"
	"sort"
)

func minMovesToSeat(seats, students []int) (ans int) {
	sort.Ints(seats)
	sort.Ints(students)
	for i, p := range seats {
		ans += abs(p - students[i])
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func main() {
	seats := []int{4, 1, 5, 9}
	students := []int{1, 3, 2, 6}
	fmt.Println(minMovesToSeat(seats, students))
}
