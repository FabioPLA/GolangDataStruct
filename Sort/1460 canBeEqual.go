package main

import (
	"fmt"
	"reflect"
	"sort"
)

func canBeEqual(target []int, arr []int) bool {
	sort.Ints(target)
	sort.Ints(arr)
	return reflect.DeepEqual(target, arr)
}
func main() {
	target := []int{1, 2, 3, 4}
	arr := []int{2, 4, 1, 3}
	fmt.Println(canBeEqual(target, arr))
}
