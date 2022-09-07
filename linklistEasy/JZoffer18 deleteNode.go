package main

func deleteNodeII(head *ListNode, val int) *ListNode {
	pre := head
	if pre.Next == nil {
		return nil
	}
	for pre.Next != nil {
		if pre.Next.Val != val {
			pre = pre.Next
		} else {
			pre.Next = pre.Next.Next
		}
	}
	return head
}
func main() {

}
