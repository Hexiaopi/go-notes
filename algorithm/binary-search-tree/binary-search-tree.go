package binary_search_tree

import (
	"fmt"
)

type Node struct {
	Left  *Node
	Data  int
	Right *Node
}

func (n *Node) Add(value int) {
	if value > n.Data {
		if n.Right == nil {
			n.Right = &Node{
				Left:  nil,
				Data:  value,
				Right: nil,
			}
		} else {
			n.Right.Add(value)
		}
	} else if value < n.Data {
		if n.Left == nil {
			n.Left = &Node{
				Left:  nil,
				Data:  value,
				Right: nil,
			}
		} else {
			n.Left.Add(value)
		}
	} else {
		fmt.Println(value, " already exist")
	}
}

func (n *Node) Find(value int) *Node {
	if n == nil {
		return nil
	}
	if n.Data == value {
		return n
	} else if n.Data > value {
		return n.Left.Find(value)
	} else {
		return n.Right.Find(value)
	}
}

func (n *Node) FindParent(value int) *Node {
	if n == nil {
		return nil
	}
	if n.Left != nil && value < n.Data {
		if n.Left.Data == value {
			return n
		} else {
			return n.Left.FindParent(value)
		}
	}
	if n.Right != nil && value > n.Data {
		if n.Right.Data == value {
			return n
		} else {
			return n.Right.FindParent(value)
		}
	}
	return nil
}

func (n *Node) Print() {
	if n != nil {
		fmt.Printf("%v", n.Data)
		if n.Left != nil || n.Right != nil {
			fmt.Printf("(")
			n.Left.Print()
			if n.Right != nil {
				fmt.Printf(",")
			}
			n.Right.Print()
			fmt.Printf(")")
		}
	}
}

// 太复杂了
func (n *Node) Delete(value int) *Node {
	node := n.Find(value)
	if node == nil {
		return n
	}
	parent := n.FindParent(value)
	if parent == nil { // 如果是根节点
		if node.Left == nil && node.Right == nil { // 根节点没有子树
			node = nil
			return nil
		}
		if node.Left != nil && node.Right != nil { // 根节点有两个子树
			// 从右子树找一个最小的值替换该节点的值
			minNode := node.Right
			for minNode.Left != nil {
				minNode = minNode.Left
			}
			node.Data = minNode.Data
			minNode.Delete(minNode.Data)
			return n
		}
		// 根节点节点有一个子树
		if node.Left != nil { // 有左子树
			node = node.Left
		} else if node.Right != nil { //有右子树
			node = node.Right
		}
		return node
	}
	// 不是根节点
	if node.Left == nil && node.Right == nil { // 删除节点没有子树
		if parent.Left != nil && value == parent.Left.Data { // 删除节点是父节点的左子树
			parent.Left = nil
		} else { // // 删除节点是父节点的右子树
			parent.Right = nil
		}
		return n
	}
	// 删除节点有左右子树
	if node.Left != nil && node.Right != nil {
		// 从右子树找一个最小的值替换该节点的值
		minNode := node.Right
		for minNode.Left != nil {
			minNode = minNode.Left
		}
		node.Data = minNode.Data
		minNode.Delete(minNode.Data)
		return n
	}
	// 删除节点有一个子树
	if node.Left != nil { //有左子树
		if parent.Left != nil && parent.Left.Data == value { // 删除节点是父节点的左子树
			parent.Left = node.Left
		} else { // 删除节点是父节点的右子树
			parent.Right = node.Left
		}
	} else {                                                 // 有右子树
		if parent.Left != nil && parent.Left.Data == value { // 删除节点是父节点的左子树
			parent.Left = node.Right
		} else { // 删除节点是父节点的右子树
			parent.Right = node.Right
		}
	}
	return n
}

// 递归版本
func (n *Node) DeleteNodeRecursion(key int) *Node {
	if n == nil {
		return nil
	}
	// 删除的节点在子树上
	if key < n.Data {
		n.Left = n.Left.DeleteNodeRecursion(key)
		return n
	}
	if key > n.Data {
		n.Right = n.Right.DeleteNodeRecursion(key)
		return n
	}
	// 删除的节点为当前节点
	if n.Right == nil { // 只有一个左子树
		return n.Left
	}
	if n.Left == nil { // 只有一个右子树
		return n.Right
	}
	// 有两个子树
	// 从右子树选择最小的node替换
	minNode := n.Right
	for minNode.Left != nil {
		minNode = minNode.Left
	}
	// 替换当前节点的值并删除右子树最小node
	n.Data = minNode.Data
	n.Right = n.Right.deleteMinNode()
	return n
}

// 递归删除节点
func (n *Node) deleteMinNode() *Node {
	if n.Left == nil { //左子树为空，则删除当前节点
		pRight := n.Right
		n.Right = nil
		return pRight
	}
	// 左子树不为空，则继续删除左子树
	n.Left = n.Left.deleteMinNode()
	return n
}

// 迭代版本
func (n *Node) DeleteNodeIteration(key int) *Node {
	// 特殊情况处理
	if n == nil {
		return n
	}
	cur := n
	var pre *Node
	for cur != nil {
		if cur.Data == key {
			break
		}
		pre = cur
		if cur.Data > key {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	if pre == nil {
		return cur.deleteOneNode()
	}
	// pre 要知道是删除左孩子还有右孩子
	if pre.Left != nil && pre.Left.Data == key {
		pre.Left = cur.deleteOneNode()
	}
	if pre.Right != nil && pre.Right.Data == key {
		pre.Right = cur.deleteOneNode()
	}
	return n
}

// 删除一个具体节点
func (n *Node) deleteOneNode() *Node {
	if n == nil {
		return n
	}
	if n.Right == nil {
		return n.Left
	}
	cur := n.Right
	for cur.Left != nil {
		cur = cur.Left
	}
	cur.Left = n.Left
	return n.Right
}
