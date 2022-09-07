package main

func kthToLast(head *ListNode, k int) int {
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}
	kth := head
	for ; n > k; n-- {
		kth = kth.Next
	}
	return kth.Val
}
func main() {

}
