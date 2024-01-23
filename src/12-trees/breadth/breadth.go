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

func Bff(head *BinaryNode, needle interface{}) bool {
	q := queues.NewQueue()
	q.Enqueue(head)
	for q.Length > 0 {
		q.Enqueue(q.Head.Value.(*BinaryNode).Left)
		q.Enqueue(q.Head.Value.(*BinaryNode).Right)
		q.Dequeue()
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
	fmt.Println(Bff(&root, 3))
}
