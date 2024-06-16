package add_two_numbers

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func Solve(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{Val: 0}
	current := result
	carry := 0

	for {
		if l1 != nil {
			current.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			current.Val += l2.Val
			l2 = l2.Next
		}
		if carry > 0 {
			current.Val += carry
			carry = 0
		}
		if current.Val > 9 {
			carry = 1
			current.Val = current.Val % 10
		}
		if l1 != nil || l2 != nil || carry > 0 {
			current.Next = &ListNode{Val: 0}
			current = current.Next
		} else {
			break
		}
	}

	return result
}
