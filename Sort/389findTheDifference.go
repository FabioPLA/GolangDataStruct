package main

import "fmt"

func findTheDifference(s string, t string) byte {
	var differ byte
	for i := range s {
		differ ^= s[i] ^ t[i]
	}
	return differ ^ t[len(t)-1]
}
func main() {
	s, t := "abcd", "abcde"
	fmt.Printf("%c", findTheDifference(s, t))
}
