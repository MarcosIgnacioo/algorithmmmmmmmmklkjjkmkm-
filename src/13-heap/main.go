package main

import (
	"fmt"

	"marcos.com/algorithms/src/13-heap/minHeap"
)

func main() {
	mn := minHeap.NewMinHeap()
	mn.Insert(4)
	mn.Insert(16)
	mn.Insert(19)
	mn.Insert(20)
	mn.Insert(21)
	mn.Insert(24)
	mn.Insert(26)
	mn.Insert(28)
	fmt.Println(mn)
	mn.Delete(3)
	mn.Delete(4)
	fmt.Println(mn)
}
