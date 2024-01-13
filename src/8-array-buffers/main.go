package main

import (
	"fmt"
	"reflect"
)

type ArrayBuffer struct {
	Array    []interface{}
	Length   uint
	Capacity uint
	Type     reflect.Type
	Head     *interface{}
	Tail     *interface{}
}

func NewArrayBuffer(c uint) ArrayBuffer {
	return ArrayBuffer{
		Array:    make([]interface{}, c),
		Length:   0,
		Capacity: c,
		Type:     nil,
		Head:     nil,
	}
}

func (ab *ArrayBuffer) Push(item interface{}) {
	array := ab.Array
	length := &ab.Length
	capacity := &ab.Capacity
	arrayType := &ab.Type
	head := &ab.Head

	if *arrayType == nil {
		*arrayType = reflect.TypeOf(item)
		*head = &array[0]
	}

	if *arrayType != reflect.TypeOf(item) {
		return
	}

	if *length+uint(1) > *capacity {
		newArray := make([]interface{}, *capacity*2)
		newArray = appendGamer(newArray, array)
		array = newArray
	}

	if *length > 0 {
		for i := *length; i > 0; i-- {
			array[i] = array[i-1]
		}
		array[0] = nil
	}
	*length++
	tmp := *head
	*tmp = item
	fmt.Println(*tmp)
}

func (ab *ArrayBuffer) Pop() {
	if ab.Length > 0 {
		ab.Length--
	}
}

func appendGamer(array []interface{}, secondArray []interface{}) []interface{} {
	for i := 0; i < len(secondArray); i++ {
		array[i] = secondArray[i]
	}
	return array
}

func main() {
	asdf := NewArrayBuffer(3)
	asdf.Push(453)
	asdf.Push(34)
	asdf.Push(45)
	asdf.Push(46)
}
