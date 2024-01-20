package inorder

import (
	"fmt"

	"marcos.com/algorithms/src/12-trees/arraylist"
	"marcos.com/algorithms/src/12-trees/stack"
)

type BinaryNode struct {
	Value interface{}
	Left  *BinaryNode
	Right *BinaryNode
}

func TestInorder() {
	// Construye un árbol de ejemplo
	//        1
	//       / \
	//      2   3
	//     / \
	//    4   5
	root := NewBinaryNode(1)
	root.Left = &BinaryNode{Value: 2, Left: &BinaryNode{Value: 4}, Right: &BinaryNode{Value: 5}}
	root.Right = &BinaryNode{Value: 3}

	// Calcula el inorder esperado para el árbol de ejemplo
	expectedResult := arraylist.NewArrayList(12)
	expectedResult.Push(4)
	expectedResult.Push(2)
	expectedResult.Push(5)
	expectedResult.Push(1)
	expectedResult.Push(3)

	// Llama a la función que estás probando
	result := Inorder(&root)
	fmt.Println(result.String())
}

func NewBinaryNode(value interface{}) BinaryNode {
	return BinaryNode{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

func Walk(current *BinaryNode, path *stack.Stack) *stack.Stack {
	// Nuestro basecase es que el nodo en el que estemos sea nulo
	if current == nil {
		return path
	}

	// Como esto es inorden la visita al nodo (la accion que queramos hacer con el en este caso insertarlo a nuestra lista) se hace entre moverse a la izquierda y derecha. Porque vamos a la izquierda hasta que no haya ningun valor, entonces pusheamos esa posicion a la lista y luego vamos a la derecha
	// recurse
	Walk(current.Left, path)
	path.Push(current.Value)
	Walk(current.Right, path)

	// Retornamos ya nuestro path cuando acabemos de recursar
	// post
	return path
}

func Inorder(head *BinaryNode) *stack.Stack {
	path := stack.Factory()
	return Walk(head, &path)
}
