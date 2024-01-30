package main

import (
	"fmt"

	"marcos.com/algorithms/src/12-trees/arraylist"
	"marcos.com/algorithms/src/14-graphs/graphs"
)

func main() {
	arr := arraylist.NewArrayList(10)
	graphs.Fill(&arr, 2)
	fmt.Println(arr.String())
}
