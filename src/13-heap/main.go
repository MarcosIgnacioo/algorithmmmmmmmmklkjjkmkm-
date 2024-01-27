package main

import (
	"fmt"

	"marcos.com/algorithms/src/13-heap/minHeap"
)

func main() {
	mn := minHeap.NewMinHeap()
	mn.Insert(2)
	mn.Insert(3)
	mn.Insert(1)
	mn.Delete(1)
	fmt.Println(mn)
}
