package main

func reversePrint(head *ListNode) []int {
	tmp := []int{}
	res := []int{}
	if head == nil {
		return nil
	}
	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}
	for len(tmp) > 0 {
		res = append(res, tmp[len(tmp)-1])
		tmp = tmp[:len(tmp)-1]
	}
	return res
}
func main() {

}
