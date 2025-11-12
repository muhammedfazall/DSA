package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) AddToFront(data int) {
	newNode := &Node{data: data}
	newNode.next = list.head
	list.head = newNode
}

func (list *LinkedList) AddToEnd(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (list *LinkedList) Display() {
	if list.head == nil {
		fmt.Println("list is empty")
	}

	current := list.head

	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}

	fmt.Println("nil")
}

func main() {
	list := LinkedList{}

	list.AddToEnd(10)
	list.AddToEnd(50)
	list.AddToEnd(60)
	list.AddToFront(70)

	list.Display()
}
