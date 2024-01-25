package main

import (
	"fmt"

	"marcos.com/algorithms/src/12-trees/bn"
	"marcos.com/algorithms/src/12-trees/dff"
)

func main() {
	//         15
	//       /   \
	//      5    17
	//     / \   |
	//    3   8 16
	tree := bn.NewTree()
	// dff.Insert(&tree, 18)
	fmt.Println(tree.Right.Left)
	dff.Delete(&tree, 22, nil)
	fmt.Println(tree.Right.Left)
	fmt.Println(tree.Right.Left.Right)
}
