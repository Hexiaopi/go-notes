package avl

type AVLNode struct {
	Value  int
	Height int
	Left   *AVLNode
	Right  *AVLNode
}

func (n *AVLNode) Find(value int) *AVLNode {
	if n == nil {
		return nil
	}
	if n.Value == value {
		return n
	} else if n.Value > value {
		return n.Left.Find(value)
	} else {
		return n.Right.Find(value)
	}
}

// 左旋转
func (n *AVLNode) LeftRotate() *AVLNode {
	headNode := n.Right
	n.Right = headNode.Left
	headNode.Left = n
	n.Height = max(n.Left.Height, n.Right.Height) + 1
	headNode.Height = max(headNode.Left.Height, headNode.Right.Height) + 1
	return headNode
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
