package doublyLinkedList

import "fmt"

type Node struct {
	Data interface{}
	Prev *Node
	Next *Node
}

func CreateDoublyLinkedList(data []interface{}) *Node {
	var head, tail *Node
	for _, value := range data {
		newNode := &Node{Data: value}
		if head == nil {
			head = newNode
			tail = newNode
		} else {
			tail.Next = newNode
			newNode.Prev = tail
			tail = newNode
		}
	}
	return head
}

func Display(head *Node) {
	curr := head
	for curr != nil {
		fmt.Print(curr.Data, " <-> ")
		curr = curr.Next
	}
	fmt.Print("nil")
	fmt.Println()
}

//Use it
