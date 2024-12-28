package linkedlist

import (
	"fmt"
)

func (ll *LinkedList) Insert(value int) {
	newNode := &Node{Value: value}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

func (ll *LinkedList) Print() {
	current := ll.Head
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}
	fmt.Println()
}

func (ll *LinkedList) Search(value int) int {
	current := ll.Head
	index := 0
	for current != nil {
		if current.Value == value {
			return index
		}
		current = current.Next
		index++
	}
	return -1
}

func (ll *LinkedList) Delete(value int) *Node {
	if ll.Head == nil {
		return nil
	}

	if ll.Head.Value == value {
		ll.Head = ll.Head.Next
		return ll.Head
	}

	current := ll.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
		}
		current = current.Next
	}
	return ll.Head
}

func (ll *LinkedList) Reverse() {
	var prev *Node
	current := ll.Head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	ll.Head = prev
}
