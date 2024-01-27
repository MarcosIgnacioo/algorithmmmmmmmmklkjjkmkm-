package minHeap

import (
	"fmt"
	"math"

	"marcos.com/algorithms/src/12-trees/arraylist"
)

type MinHeap struct {
	Length int
	Data   arraylist.ArrayList
}

func (mh MinHeap) String() string {
	return mh.Data.String()
}

func (mh *MinHeap) BubbleUp(i int) {
	fIdx := GetFather(i)
	father := mh.Data.ArrayList[fIdx]
	node := mh.Data.ArrayList[i]
	if node.(int) < father.(int) {
		mh.Data.ArrayList[fIdx] = node
		mh.Data.ArrayList[i] = father
		mh.BubbleUp(fIdx)
	} else {
		return
	}
}

func (mh *MinHeap) BubbleDown(i int) {
	var smallestChild interface{}
	node := mh.Data.ArrayList[i].(int)
	l, r := GetChildren(i)
	sIdx := int(math.Min(float64(l), float64(r)))

	smallestChild = mh.Data.ArrayList[sIdx]
	if smallestChild != nil {

		if smallestChild.(int) < node {
			mh.Data.ArrayList[sIdx] = node
			mh.Data.ArrayList[i] = smallestChild
			mh.BubbleDown(sIdx)
		}
	}
	return
}

func (mh *MinHeap) Delete(i int) {
	index := mh.Data.Length - 1
	if i == int(index) {
		mh.Data.Dequeue()
		mh.Length--
		return
	}
	deleted := mh.Data.ArrayList[index]
	mh.Data.ArrayList[i] = deleted
	mh.Data.ArrayList[index] = nil
	mh.Data.Length--
	mh.Length--
	fmt.Println("i: ", i)
	mh.BubbleDown(i)
}

func (mh *MinHeap) Insert(v interface{}) {
	mh.Data.Enqueue(v)
	index := mh.Data.Length - 1
	mh.Length++
	mh.BubbleUp(int(index))
}

func GetFather(i int) int {
	return (i - 1) / 2
}

func GetChildren(i int) (l int, r int) {
	l = (2 * i) + 1
	r = (2 * i) + 2
	return
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		Length: 0,
		Data:   arraylist.NewArrayList(5),
	}
}
