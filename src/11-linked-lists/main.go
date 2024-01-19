package main

import (
	"errors"
	"fmt"
)

func main() {
	ll := NewLinkedList()
	ll.Prepend(4)
	ll.Prepend(9)
	ll.Prepend(11)
	ll.Prepend(13)
	ll.InsertAt(5, 1)
	curr := ll.Head
	for i := 0; i < ll.GetLength(); i++ {
		fmt.Println(curr.Value)
		curr = curr.Next
	}
	fmt.Println("/////////////")
	fmt.Println("Tail antes:", ll.Tail.Value)
	ll.Remove(4)
	fmt.Println("Tail despues:", ll.Tail.Value)
	curr = ll.Head
	for i := 0; i < ll.GetLength(); i++ {
		fmt.Println(curr.Value)
		curr = curr.Next
	}
}

type Node struct {
	Value interface{}
	Prev  *Node
	Next  *Node
}

func NewNode(value interface{}) *Node {
	return &Node{
		Value: value,
		Prev:  nil,
		Next:  nil,
	}
}

type LinkedList struct {
	Length int
	Head   *Node
	Tail   *Node
}

func NewLinkedList() LinkedList {
	return LinkedList{
		Length: 0,
		Head:   nil,
		Tail:   nil,
	}
}

func (L *LinkedList) GetLength() int {
	return L.Length
}

func (L *LinkedList) Get(idx int) (interface{}, error) {
	curr := L.Head
	for i := 0; i < L.GetLength(); i++ {
		if i == idx {
			return curr.Value, nil
		}
		curr = curr.Next
	}
	return nil, errors.New("Not found")
}

func (L *LinkedList) InsertAt(item interface{}, idx int) {
	if idx > L.GetLength() {
		fmt.Println("leng")
		return
	}

	if idx == L.GetLength() {
		fmt.Println("appe")
		L.Append(item)
		return
	} else if idx == L.GetLength() {
		fmt.Println("prepapen")
		L.Prepend(item)
		return
	}
	// Bookeeping aqui porque preappend y append lo hacen ellos solos
	L.Length++
	curr := L.Head
	for i := 0; i < idx; i++ {
		curr = curr.Next
	}
	node := NewNode(item)
	prev := curr.Prev

	prev.Next = node
	node.Next = curr
	node.Prev = prev
	curr.Prev = node
}

func (L *LinkedList) Remove(item interface{}) (interface{}, error) {
	curr := *L.Head
	for i := 0; i < L.GetLength(); i++ {
		if curr.Value == item {
			L.Length--
			if L.Length == 0 {
				L.Head = nil
				L.Tail = nil
				return item, nil
			}
			prev := curr.Prev
			next := curr.Next
			head := *L.Head
			tail := *L.Tail
			if prev != nil {
				prev.Next = next
			}
			if next != nil {
				next.Prev = prev
			}

			if curr == head {
				L.Head = curr.Next
			}

			if curr == tail {
				L.Tail = curr.Prev
			}

			curr.Next = nil
			curr.Prev = nil
			return item, nil
		}
		curr = *curr.Next
	}
	return nil, errors.New("Item not found")
}

func (L *LinkedList) RemoveAt(idx int) (interface{}, error) {
	curr := L.Head
	for i := 0; i < L.GetLength(); i++ {
		if i == idx {
			prev := curr.Prev
			next := curr.Next
			prev.Next = next
			next.Prev = prev
			return curr, nil
		}
		curr = curr.Next
	}
	return nil, errors.New("Item not found")

}

func (L *LinkedList) Append(item interface{}) {
	node := NewNode(item)
	L.Length++

	if L.Tail == nil {
		L.Head = node
		L.Tail = node
		return
	}

	tail := L.Tail
	node.Prev = tail
	tail.Prev = node
	L.Tail = node
}

func (L *LinkedList) Prepend(item interface{}) {
	// BOOKEEPING
	node := NewNode(item)
	L.Length++
	if L.Head == nil {
		L.Head = node
		L.Tail = node
		return
	}
	head := L.Head
	node.Next = head
	head.Prev = node
	L.Head = node
}
