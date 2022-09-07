package main

func getDecimalValue(head *ListNode) int {
	ans := 0
	for head != nil {
		ans = 2*ans + head.Val
		head = head.Next
	}
	return ans
}
func main() {

}
