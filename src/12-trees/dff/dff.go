package dff

import (
	"marcos.com/algorithms/src/12-trees/bn"
)

// Para que esto funcione los arboles en los que busquemos deben de estar odernados de la siguiente manera
// Es decir menores a la izquierda y mayores a la derecha
//        15
//       / \
//      5  17
//     / \
//    3   8
func Find(n *bn.BinaryNode, v int) bool {

	if n == nil {
		return false
	}

	if v == n.Value.(int) {
		return true
	}

	if n.Value.(int) < v {
		return Find(n.Right, v)
	}
	return Find(n.Left, v)
}

func Insert(n *bn.BinaryNode, v int) {

	if n.Left == nil && v <= n.Value.(int) {
		asdf := bn.NewBinaryNode(v)
		n.Left = &asdf
		return
	}

	if n.Right == nil && v > n.Value.(int) {
		asdf := bn.NewBinaryNode(v)
		n.Right = &asdf
		return
	}

	if n.Value.(int) < v {
		Insert(n.Right, v)
	}
	Insert(n.Left, v)
}

func InsertThePrimeagen(n *bn.BinaryNode, v int) {
	if v > n.Value.(int) {
		if n.Right == nil {
			buffer := bn.NewBinaryNode(v)
			n.Right = &buffer
			return
		}
		InsertThePrimeagen(n.Right, v)
	} else if v < n.Value.(int) {
		if n.Left == nil {
			buffer := bn.NewBinaryNode(v)
			n.Left = &buffer
			return
		}
		InsertThePrimeagen(n.Left, v)
	}
}

func Delete(n *bn.BinaryNode, v int, p *bn.BinaryNode) {
	if p == nil {
		p = n
	}
	if n == nil {
		return
	}
	if v == n.Value.(int) {
		if n.Left != nil && n.Right != nil {
			if v > p.Value.(int) {
				smallestGreatest := SmallestGreaterSide(n)
				n.Value = smallestGreatest.Value
				smallestGreatest = nil
			} else {
				greatestSmallest := GreatestSmallerSide(n)
				n.Value = greatestSmallest.Value
				greatestSmallest = nil
			}
			return
		} else if n.Left != nil && n.Right == nil {
			p.Left = n.Left
			return
		} else if n.Right != nil && n.Left == nil {
			p.Right = n.Right
			return
		} else {
			Remove(p, v)
			return
		}
	}
	p = n
	if n.Value.(int) < v {
		Delete(n.Right, v, p)
	}
	Delete(n.Left, v, p)
}

func GreatestSmallerSide(n *bn.BinaryNode) *bn.BinaryNode {
	if n.Right != nil {
		return GreatestSmallerSide(n.Right)
	}
	return n
}

func SmallestGreaterSide(n *bn.BinaryNode) *bn.BinaryNode {
	if n.Left != nil {
		return SmallestGreaterSide(n.Left)
	}
	return n
}

func Remove(n *bn.BinaryNode, v int) {
	if n.Left.Value.(int) == v {
		n.Left = nil
	} else if n.Right.Value.(int) == v {
		n.Right = nil
	}
}
func ParentToChild(p *bn.BinaryNode, c *bn.BinaryNode, v int) {
	if v > p.Value.(int) {
		p.Right = c
	} else {
		p.Left = c
	}
}
