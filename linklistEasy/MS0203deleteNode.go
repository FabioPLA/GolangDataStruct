package main

func deleteNodeIII(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
func main() {

}
