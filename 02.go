package main

import "fmt"

// Create a struct for a linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Create the anchor variables:
	// 	* result: the linked list with the result digits
	var result *ListNode
	// 	* tail: the tail of the result list
	var tail *ListNode
	// 	* carry: the carry of the addition (vai a um)
	carry := 0

	// Iterate while there are nodes in l1 or l2, or carry != 0
	for l1 != nil || l2 != nil || carry != 0 {
		// Get the values of the current nodes, if they exist
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			// Move to the next node
			l1 = l1.Next
		}
		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			// Move to the next node
			l2 = l2.Next
		}
		// Sum the node values with the carry
		sum := val1 + val2 + carry
		// Update the carry | It's gonna be 1 if sum >= 10, else 0
		carry = sum / 10
		// Get the new digit | It's gonna be the rest of the sum | sum % 10
		newDigit := sum % 10
		// Create the new node
		newNode := &ListNode{Val: newDigit}
		// If the result is nil, create it
		if result == nil {
			result = newNode
			tail = newNode
		} else { // Otherwise, append the new node to the tail
			tail.Next = newNode
			tail = newNode
		}
	}
	return result
}

func main2() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	fmt.Println(addTwoNumbers(l1, l2))
}
