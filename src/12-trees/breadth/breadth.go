package breadth

import (
	"fmt"

	"marcos.com/algorithms/src/12-trees/queues"
)

type BinaryNode struct {
	Value interface{}
	Left  *BinaryNode
	Right *BinaryNode
}

func NewBinaryNode(value interface{}) BinaryNode {
	return BinaryNode{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

//        1 <--- lo metemos a la cola
//       / \
//      2   3
//     / \
//    4   5

// Metemos a la cola  a nuestra root
// Iniciamos el ciclo
//+-->Sacamos a nuestra head actual (Length disminuye)
//\  Vemos si sus valores de izquierda o derecha no son nil
//+- Si no son nil (es decir que existen), los metemos a la cola (Length aumenta)

// Si el valor de head que sacamos vale lo mismo que el needle devolvemos
// Si el ciclo se termina no lo encontro por lo que no existe en el arbol

func Bff(head *BinaryNode, needle interface{}) bool {
	// Creamos nuestra cola
	q := queues.NewQueue()
	q.Enqueue(head)

	for q.Length > 0 {
		// Creamos nuestra variable que apunta a head
		curr := q.Head
		// Sacamos a nuestra head actual y la guardamos en una variable
		// Si el valor que buscamos es igual al que sacamos retornamos como true
		pop := q.Dequeue().(*BinaryNode).Value
		if pop == needle {
			return true
		}
		// Aqui lo que hacemos es meternos al valor de nuestro current que era la head, checamos, como el valor de Value en el struct original es una interface {} tenemos que indicar el tipo de dato que esperamos que tenga nuestro value, en este caso pues el value va a contener a un tipo de un nodo binario. Checamos que si la izquierda de este no es nil, es decir, que tenga un valor, si hay un valor lo metemos a la cola, si no pues no, y pues es la misma tanto para la derecha ocmo la izquierda
		if curr.Value.(*BinaryNode).Left != nil {
			q.Enqueue(curr.Value.(*BinaryNode).Left)
		}
		if curr.Value.(*BinaryNode).Right != nil {
			q.Enqueue(curr.Value.(*BinaryNode).Right)
		}
	}
	return false
}

func TestBff() {
	// Create a sample binary tree for testing
	// Example:
	//        1
	//       / \
	//      2   3
	//     / \
	//    4   5
	root := NewBinaryNode(1)
	root.Left = &BinaryNode{Value: 2, Left: &BinaryNode{Value: 4}, Right: &BinaryNode{Value: 5}}
	root.Right = &BinaryNode{Value: 3}
	fmt.Println(Bff(&root, 6))
}

