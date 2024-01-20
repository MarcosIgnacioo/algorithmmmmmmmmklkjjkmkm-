package preorder

import (
	"fmt"

	"marcos.com/algorithms/src/12-trees/stack"
)

type BinaryNode struct {
	Value interface{}
	Left  *BinaryNode
	Right *BinaryNode
}

func TestPreorder() {
	// Construye un árbol de ejemplo
	//        1
	//       / \
	//      2   3
	//     / \
	//    4   5
	root := NewBinaryNode(1)
	root.Left = &BinaryNode{Value: 2, Left: &BinaryNode{Value: 4}, Right: &BinaryNode{Value: 5}}
	root.Right = &BinaryNode{Value: 3}

	result := Preorder(&root)
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

	// En el caso de que no sea nulo pues pusheamos el valor a nuestra lista
	// pre
	path.Push(current.Value)

	// Ahora recursamos yendo a la izquierda y a la derecha de nuestro nodo, como ponemos primero el de la izquierda pues el path va a poner primero lo de la izquierda en nuestro y luego lo de right por lo que esto se puede hacer asi sin pedillos
	// recurse
	Walk(current.Left, path)
	Walk(current.Right, path)

	// Retornamos ya nuestro path cuando acabemos de recursar
	// post
	return path
}

func Preorder(head *BinaryNode) *stack.Stack {
	path := stack.Factory()
	return Walk(head, &path)
}
