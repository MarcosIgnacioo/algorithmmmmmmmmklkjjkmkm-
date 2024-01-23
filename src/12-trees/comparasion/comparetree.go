package comparasion

import (
	"fmt"
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

func primeagensVersion(a *BinaryNode, b *BinaryNode) bool {
	// structural check
	// Checamos si ambos son nil si son nil pues en ambos ya no sigue habiendo mas nodos hijos
	if a == nil && b == nil {
		return true
	}
	// structural check
	// Checamos si uno de los dos es nil, en caso de que solamente uno de los dos sea nil significa que otro de los nodos sigue teniendo hijos por lo que no son iguales estructuralmente
	if a == nil || b == nil {
		return false
	}
	// value check
	// Si estructuralmente son iguales, es decir que aun tienen mas hijos pues checamos los valores, y si no son iguales retornamos false
	if a.Value != b.Value {
		return false
	}
	// Al final lo que retornamos es la comparacion de las dos ramas, derecha e izquierda, si una de las dos es diferente pues sera false, e ira escalando hacia arriba en el stack de las funciones en memoria
	return primeagensVersion(a.Left, b.Left) && primeagensVersion(a.Right, b.Right)
}

func Walk(firstHead *BinaryNode, secondHead *BinaryNode, res *bool) bool {
	// Nuestro basecase es que el nodo en el que estemos sea nulo

	if firstHead == nil || secondHead == nil {
		if firstHead == nil && secondHead != nil {
			*res = false
		}

		if firstHead != nil && secondHead == nil {
			*res = false
		}
		return *res
	}

	if firstHead.Value != secondHead.Value {
		*res = false
	}

	firstLeft := firstHead.Left
	secondLeft := secondHead.Left

	firstRight := firstHead.Right
	secondRight := secondHead.Right

	Walk(firstLeft, secondLeft, res)
	Walk(firstRight, secondRight, res)

	return *res
}

func Compare(head *BinaryNode, secondHead *BinaryNode) bool {
	// res := true
	// return Walk(head, secondHead, &res)
	return primeagensVersion(head, secondHead)
}

//        1 <--- lo metemos a la cola
//       / \
//      2   3
//     / \
//    4   5

// Metemos a la cola  a nuestra root
// Iniciamos el ciclo

//+-->Sacamos a nuestra head actual (Length disminuye)
//\   Vemos si sus valores de izquierda o derecha no son nil
//+---Si no son nil (es decir que existen), los metemos a la cola (Length aumenta)

// Si el valor de head que sacamos vale lo mismo que el needle devolvemos
// Si el ciclo se termina no lo encontro por lo que no existe en el arbol

func TestCompare() {
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

	root2 := NewBinaryNode(1)
	root2.Left = &BinaryNode{Value: 2, Left: &BinaryNode{Value: 4}, Right: &BinaryNode{Value: 5}}
	root2.Right = &BinaryNode{Value: 3}

	fmt.Println(Compare(&root, &root2))
}
