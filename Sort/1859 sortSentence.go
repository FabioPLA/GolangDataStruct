package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortSentence(s string) string {
	strs := strings.Split(s, " ")
	sort.Slice(strs, func(i, j int) bool {
		return strs[i][len(strs[i])-1] < strs[j][len(strs[j])-1]
	})
	for i := range strs {
		strs[i] = strs[i][:len(strs[i])-1]
	}
	return strings.Join(strs, " ")
}
func main() {
	s := "is2 sentence4 This1 a3"
	fmt.Println(sortSentence(s))
}
