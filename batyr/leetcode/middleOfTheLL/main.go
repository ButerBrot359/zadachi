package main

func middleNode(head *ListNode) *ListNode {
	needSize := listCount(head) / 2
	for needSize != 0 {
		head = head.Next
		needSize--
	}
	return head

}
func listCount(head *ListNode) int {
	count := 0
	for head != nil {
		head = head.Next
		count++
	}
	return count
}
