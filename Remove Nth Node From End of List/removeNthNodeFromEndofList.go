package main

func main() {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Val: 0, Next: head}
	delNode := dummyNode
	for i := 0; i < n; i++ {
		head = head.Next
	}
	for head != nil {
		delNode = delNode.Next
		head = head.Next
	}

	delNode.Next = delNode.Next.Next

	return dummyNode.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
