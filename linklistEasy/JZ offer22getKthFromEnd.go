package main

//func getKthFromEnd(head *ListNode, k int) *ListNode {
//	tmp := []int{}
//	res := []int{}
//	for head != nil {
//		tmp = append(tmp, head.Val)
//		head = head.Next
//	}
//	for i := 0; i <= k-1; i++ {
//		res = append(res, tmp[i])
//	}
//	return res
//}
func getKthFromEnd(head *ListNode, k int) (kth *ListNode) {
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}
	for kth = head; n > k; n-- {
		kth = kth.Next
	}
	return
}
func main() {

}
