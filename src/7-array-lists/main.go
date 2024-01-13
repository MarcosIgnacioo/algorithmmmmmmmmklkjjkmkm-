package main

import (
	"fmt"
	"reflect"
)

type ArrayList struct {
	ArrayList []interface{}
	Length    uint
	Capacity  uint
	Type      reflect.Type
}

func NewArrayList(c uint) ArrayList {
	return ArrayList{
		ArrayList: make([]interface{}, c),
		Length:    0,
		Capacity:  c,
	}
}

func (arrayList *ArrayList) Enqueue(item interface{}) interface{} {

	if arrayList.Type == nil {
		arrayList.Type = reflect.TypeOf(item)
	}

	if arrayList.Type != reflect.TypeOf(item) {
		return nil
	}

	if arrayList.Length+1 > arrayList.Capacity {
		newArray := make([]interface{}, arrayList.Capacity*2)
		newArray = appendGamer(newArray, arrayList.ArrayList)
		arrayList.ArrayList = newArray
	}

	arrayList.ArrayList[arrayList.Length] = item
	arrayList.Length++
	return item
}

func (arrayList *ArrayList) Push(item interface{}) interface{} {
	if arrayList.Type == nil {
		arrayList.Type = reflect.TypeOf(item)
	}

	if arrayList.Type != reflect.TypeOf(item) {
		return nil
	}

	if arrayList.Length+1 > arrayList.Capacity {
		newArray := make([]interface{}, arrayList.Capacity*2)
		newArray = appendGamer(newArray, arrayList.ArrayList)
		arrayList.ArrayList = newArray
	}

	if arrayList.Length > 0 {
		for i := arrayList.Length; i > 0; i-- {
			arrayList.ArrayList[i] = arrayList.ArrayList[i-1]
		}
	}
	arrayList.ArrayList[0] = item
	arrayList.Length++
	return item
}

func (arrayList *ArrayList) Pop() {
	arrayList.ArrayList[arrayList.Length] = nil
	if arrayList.Length > 0 {
		arrayList.Length--
	}
}

func (arrayList *ArrayList) Dequeue() interface{} {

	var dequeued interface{}
	if arrayList.Length > 0 {
		dequeued = arrayList.ArrayList[0]
		for i := 0; i < int(arrayList.Length)-1; i++ {
			arrayList.ArrayList[i] = arrayList.ArrayList[i+1]
		}
		arrayList.ArrayList[arrayList.Length-1] = nil
	}

	if arrayList.Length > 0 {
		arrayList.Length--

	}
	return dequeued
}

func appendGamer(array []interface{}, secondArray []interface{}) []interface{} {
	for i := 0; i < len(secondArray); i++ {
		array[i] = secondArray[i]
	}
	return array
}

func main() {
	// Create a new ArrayList with capacity 5
	myList := NewArrayList(5)

	// Enqueue some items
	myList.Enqueue(1)
	myList.Enqueue("two")
	myList.Enqueue(3.0)

	// Print the ArrayList
	fmt.Println("After Enqueue:")
	printArrayList(&myList)

	// Push an item to the front
	myList.Push("zero")

	// Print the ArrayList
	fmt.Println("\nAfter Push:")
	printArrayList(&myList)

	// Dequeue an item
	dequeued := myList.Dequeue()
	fmt.Printf("\nDequeued: %v\n", dequeued)

	// Print the ArrayList
	fmt.Println("After Dequeue:")
	printArrayList(&myList)

	// Pop an item
	myList.Dequeue()

	// Print the ArrayList
	fmt.Println("\nAfter Pop:")
	printArrayList(&myList)
}

func printArrayList(arrayList *ArrayList) {
	fmt.Printf("Length: %d, Capacity: %d, Type: %v\n", arrayList.Length, arrayList.Capacity, arrayList.Type)
	fmt.Printf("ArrayList: %v\n", arrayList.ArrayList)
}
