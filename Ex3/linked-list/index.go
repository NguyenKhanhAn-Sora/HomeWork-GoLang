package main

import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
}

func (ll *LinkedList[T]) Prepend(value T) {
	newNode := &Node[T]{data: value, next: ll.head}
	ll.head = newNode
}

func (ll *LinkedList[T]) Append(value T) {
	newNode := &Node[T]{data: value}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head

	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (ll *LinkedList[T]) PrintList() {
	current := ll.head

	for current != nil {
		fmt.Printf("%v ->", current.data)
		current = current.next
	}

	fmt.Println("nil")
}

func (ll *LinkedList[T]) Search(value T) bool {
	current := ll.head
	for current.next != nil {
		if current.data == value {
			return true
		}
	}

	return false

}

func (ll *LinkedList[T]) Delete(value T) {
	if ll.head == nil {
		return
	}

	if ll.head.data == value {
		ll.head = ll.head.next
	}

	current := ll.head

	for current.next != nil && current.next.data != value {
		current = current.next
	}

	if current.next != nil {
		current.next = current.next.next
	}
}

//Viết hàm Insert At
func (ll *LinkedList[T]) InsertAt(position int, value T) {
	if position == 0 {
		ll.Prepend(value)
		return
	}
	newNode := &Node[T]{data: value}
	current := ll.head

	for i := 0; i < position-1; i++ { //0,1,2,3,4
		if current == nil {
			return
		} else {
			current = current.next
		}
	}
	newNode.next = current.next
	current.next = newNode
}

func main() {
	ll := LinkedList[int]{}
	ll.Append(10)
	ll.Append(20)
	// 10 -> 20 -> nil
	ll.Prepend(5)
	// 5->10 -> 20 -> nil
	ll.PrintList()

	fmt.Println("Searching 10...", ll.Search(10))
	fmt.Println("Searching 30...", ll.Search(30))
}
