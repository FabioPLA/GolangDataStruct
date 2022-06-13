package main

import "fmt"

func checkIfExist(arr []int) bool {
	for i, k := range arr {
		for j, v := range arr {
			if i != j && k*2 == v {
				return true
			}
		}
	}
	return false
}
func main() {
	arr := []int{10, 2, 5, 3}
	fmt.Println(checkIfExist(arr))
}
