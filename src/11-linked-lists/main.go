package main

import (
	"errors"
	"fmt"
)

func main() {
	ll := NewLinkedList()
	ll.Append(4)
	ll.Append(9)
	ll.Append(11)
	fmt.Println(ll.String())
	fmt.Println(ll.String())
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

func (L *LinkedList) String() string {
	curr := L.Head
	List := "["
	for i := 0; i < L.Length; i++ {
		if i == L.Length-1 {
			List += fmt.Sprintf("%v", curr.Value)
			curr = curr.Next
			continue
		}
		if curr != nil {
			List += fmt.Sprintf("%v,", curr.Value)
			curr = curr.Next
		} else {
			return List
		}
	}
	List += "]"
	return List
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
	item, err := L.GetAt(idx)
	if err != nil {
		return nil, errors.New("Not found")
	}
	return item.Value, nil
}

func (L *LinkedList) InsertAt(item interface{}, idx int) {
	if idx > L.GetLength() {
		return
	}
	if idx == 0 {
		L.Prepend(item)
		return
	} else if idx == L.GetLength() {
		L.Append(item)
		return
	}
	// Bookeeping aqui porque preappend y append lo hacen ellos solos
	L.Length++
	curr, err := L.GetAt(idx)
	if err != nil {
		return
	}
	node := NewNode(item)
	prev := curr.Prev

	prev.Next = node
	node.Next = curr
	node.Prev = prev
	curr.Prev = node
}

func (L *LinkedList) Remove(item interface{}) (interface{}, error) {
	curr := L.Head
	for i := 0; i < L.GetLength(); i++ {
		if curr.Value == item {
			return L.DeleteNode(curr)
		}
		curr = curr.Next
	}
	return nil, errors.New("Item not found")
}

func (L *LinkedList) DeleteNode(curr *Node) (*Node, error) {
	L.Length--
	if L.Length == 0 {
		L.Head = nil
		L.Tail = nil
		return curr, nil
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

	if *curr == head {
		L.Head = curr.Next
	}

	if *curr == tail {
		L.Tail = curr.Prev
	}

	curr.Next = nil
	curr.Prev = nil
	return curr, nil
}

func (L *LinkedList) RemoveAt(idx int) (interface{}, error) {
	curr := L.Head
	for i := 0; i < L.Length; i++ {
		if i == idx {
			return L.DeleteNode(curr)
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
	tail.Next = node
	L.Tail = node
	// Podemos usar la funcion de string para debuggear para saber que elementos tiene nuestra lista por cada operacion que se hace
	fmt.Println(L.String())
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

func (L *LinkedList) GetAt(idx int) (*Node, error) {
	curr := L.Head
	if idx == 0 {
		return L.Head, nil
	} else if idx == L.Length-1 {
		return L.Tail, nil
	}
	for i := 0; i < idx; i++ {
		if curr != nil {
			curr = curr.Next
		} else {
			return nil, errors.New("Not found")
		}
	}
	return curr, nil
}
