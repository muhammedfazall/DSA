package main

import "fmt"

//structure of a Node: a single element in the linked list
type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
}

// AddToFront adds a new node at the beginning of the list
func (list *LinkedList) AddToFront(data int) {
	newNode := &Node{data: data} // creating new Node
	newNode.next = list.head     // point newNode.next to current head
	list.head = newNode          // move head to point to newNode
}

func main() {

	list := LinkedList{}

	list.AddToFront(50)

	fmt.Println("Added first node with value:", list.head.data)
	fmt.Println("next of head:", list.head.next)
}
