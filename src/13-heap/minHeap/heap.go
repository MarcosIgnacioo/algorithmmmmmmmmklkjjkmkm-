package minHeap

import (
	"fmt"
	"marcos.com/algorithms/src/12-trees/arraylist"
	"math"
)

type MinHeap struct {
	Length int
	Data   arraylist.ArrayList
}

func (mh MinHeap) String() string {
	return mh.Data.String()
}

func (mh *MinHeap) Delete(i int) {
	index := mh.Data.Length - 1
	if i == int(index) {
		mh.Data.Pop()
		mh.Length--
		return
	}
	if i > mh.Length {
		return
	}
	deleted := mh.Data.ArrayList[index]
	mh.Data.ArrayList[i] = deleted
	mh.Data.ArrayList[index] = nil
	mh.Data.Length--
	mh.Length--
	// mh.BubbleDown(i)
	mh.HeapifyDown(i)
}

func (mh *MinHeap) HeapifyDown(i int) {
	lIdx, rIdx := GetChildrenIndexes(i)
	if i > mh.Length || lIdx >= mh.Length || rIdx >= mh.Length {
		return
	}
	node := mh.Data.ArrayList[i].(int)
	lV := mh.Data.ArrayList[lIdx].(int)
	rV := mh.Data.ArrayList[rIdx].(int)
	sIdx := lIdx

	if lV > rV {
		sIdx = rIdx
		fmt.Println("wh")
	}
	smallestChild := mh.Data.ArrayList[sIdx]

	if smallestChild.(int) < node {
		mh.Swap(i, sIdx)
		mh.HeapifyDown(sIdx)
	} else {
	}
	return
}

func (mh *MinHeap) BubbleDown(i int) {

	var smallestChild interface{}
	var sIdx int
	node := mh.Data.ArrayList[i].(int)
	l, r := mh.GetChildren(i)
	fmt.Println(l, r)

	if l.(int) > mh.Length && r.(int) > mh.Length {
		return
	}
	if (l.(int) > mh.Length) != (r.(int) > mh.Length) {
		sIdx = SelectNotNil(l, r)
	}
	if l != nil && r != nil {
		fmt.Println("--------")
		fmt.Println(l, r)
		sIdx = int(math.Min(float64(l.(int)), float64(r.(int))))
		fmt.Println("//////")
		fmt.Println(sIdx)
	}
	smallestChild = mh.Data.ArrayList[sIdx]
	if smallestChild != nil {
		if smallestChild.(int) < node {
			// mh.Data.ArrayList[sIdx] = node
			// mh.Data.ArrayList[i] = smallestChild
			mh.Swap(sIdx, i)
			mh.BubbleDown(sIdx)
		}
		return
	}

}

func SelectNotNil(l interface{}, r interface{}) int {
	if l == nil {
		return r.(int)
	}
	return l.(int)
}

func (mh *MinHeap) Insert(v interface{}) {
	mh.Data.Enqueue(v)
	index := mh.Data.Length - 1
	mh.Length++
	mh.BubbleUp(int(index))
}

func (mh *MinHeap) BubbleUp(i int) {
	if i == 0 {
		return
	}
	fIdx := GetFather(i)
	father := mh.Data.ArrayList[fIdx]
	node := mh.Data.ArrayList[i]

	if node.(int) < father.(int) {
		// mh.Data.ArrayList[fIdx] = node
		// mh.Data.ArrayList[i] = father
		mh.Swap(fIdx, i)
		mh.BubbleUp(fIdx)
	} else {
		return
	}
}
func (mh *MinHeap) Swap(a int, b int) {
	tmp := mh.Data.ArrayList[a]
	mh.Data.ArrayList[a] = mh.Data.ArrayList[b]
	mh.Data.ArrayList[b] = tmp
}

func GetFather(i int) int {
	return (i - 1) / 2
}

func (mh *MinHeap) GetChildren(i int) (l interface{}, r interface{}) {
	l, r = GetChildrenIndexes(i)
	return
}
func GetChildrenIndexes(i int) (l int, r int) {
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
