package main

import (
	"fmt"

	"marcos.com/algorithms/src/12-trees/inorder"
	"marcos.com/algorithms/src/12-trees/postorder"
	"marcos.com/algorithms/src/12-trees/preorder"
)

func main() {
	fmt.Println("Preorder")
	preorder.TestPreorder()
	fmt.Println("Postorder")
	postorder.TestPostorder()
	fmt.Println("Inorder")
	inorder.TestInorder()
}
