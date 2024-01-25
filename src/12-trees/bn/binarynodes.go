package bn

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

func NewTree() (root BinaryNode) {

	//        15
	//       / \
	//      5  17
	//     / \
	//    3   8

	//           15
	//         /    \
	//        /     \
	//        5     27
	//      /  \   /       |
	//     /   \  22       31
	//    /    \ 20 25   29 36
	//   3     8
	//  /    / |
	// 2
	//
	//
	root = NewBinaryNode(15)

	root.Left = &BinaryNode{
		Value: 5,
		Left:  &BinaryNode{Value: 3},
		Right: &BinaryNode{Value: 8},
	}
	root.Left.Left.Left = &BinaryNode{Value: 2}
	root.Right = &BinaryNode{Value: 27, Left: &BinaryNode{Value: 22, Left: &BinaryNode{Value: 20}, Right: &BinaryNode{Value: 25}}, Right: &BinaryNode{Value: 31, Left: &BinaryNode{Value: 29}, Right: &BinaryNode{Value: 36}}}
	return
}
