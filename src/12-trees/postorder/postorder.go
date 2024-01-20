package postorder

import (
	"fmt"
	"marcos.com/algorithms/src/12-trees/arraylist"
)

type BinaryNode struct {
	Value interface{}
	Left  *BinaryNode
	Right *BinaryNode
}

func TestPostorder() {
	// Construye un árbol de ejemplo
	//        1
	//       / \
	//      2   3
	//     / \
	//    4   5
	root := NewBinaryNode(1)
	root.Left = &BinaryNode{Value: 2, Left: &BinaryNode{Value: 4}, Right: &BinaryNode{Value: 5}}
	root.Right = &BinaryNode{Value: 3}

	// Calcula el postorden esperado para el árbol de ejemplo
	expectedResult := arraylist.NewArrayList(12)
	expectedResult.Push(4)
	expectedResult.Push(5)
	expectedResult.Push(2)
	expectedResult.Push(3)
	expectedResult.Push(1)

	// Llama a la función que estás probando
	result := Postorder(&root)
	fmt.Println(result.ArrayList...)
}
func NewBinaryNode(value interface{}) BinaryNode {
	return BinaryNode{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

func Walk(current *BinaryNode, path *arraylist.ArrayList) *arraylist.ArrayList {
	// Nuestro basecase es que el nodo en el que estemos sea nulo
	if current == nil {
		return path
	}

	// recurse
	Walk(current.Left, path)
	Walk(current.Right, path)

	// Como esto es postorden hacemos el movimiento de izquierda y derecha primero y ya luego lo registramos en nuestra lista, es decir cuando
	// post
	path.Enqueue(current.Value)
	return path
}

func Postorder(head *BinaryNode) *arraylist.ArrayList {
	path := arraylist.NewArrayList(12)
	return Walk(head, &path)
}
