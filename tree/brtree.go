package tree

//BinaryNode  二叉树节点
type BinaryNode struct {
	Val    int
	Left   *BinaryNode
	Right  *BinaryNode
	Parent *BinaryNode
}

// Search 二叉排序树搜索
func (root *BinaryNode) Search(value int) *BinaryNode {
	node := root
	for node != nil {
		if node.Val == value {
			break
		} else if node.Val < value {
			node = node.Right
		} else {
			node = node.Left
		}
	}
	return node
}

// Add 二叉排序树插入
func (root *BinaryNode) Add(value int) bool {
	node := root
	for {
		if node.Val == value {
			return false
		} else if node.Val < value {
			if node.Right == nil {
				node.Right = &BinaryNode{
					Val:    value,
					Parent: node,
				}
				break
			} else {
				node = node.Right
			}
		} else {
			if node.Left == nil {
				node.Left = &BinaryNode{
					Val:    value,
					Parent: node,
				}
			} else {
				node = node.Left
			}
		}
	}
	return true
}

// Delete 排序二叉树删除
func (root *BinaryNode) Delete(value int) (*BinaryNode, bool) {
	node := root.Search(value)
	if node == nil {
		return root, false
	}
	return root.delete(node)
}

func (root *BinaryNode) delete(node *BinaryNode) (*BinaryNode, bool) {
	if node.Right == nil && node.Left == nil {
		if node.Parent != nil {
			if node.isLeftChild() {
				node.Parent.Left = nil
			} else {
				node.Parent.Right = nil
			}
		} else {
			root = nil
		}
		node = nil
	} else if node.Right == nil {
		if node.Parent != nil {
			if node.isLeftChild() {
				node.Parent.Left = node.Left
				node.Left.Parent = node.Parent
			} else {
				node.Parent.Right = node.Left
				node.Left.Parent = node.Parent
			}
		} else {
			root = node.Left
		}
		node = nil
	} else if node.Left == nil {
		if node.Parent != nil {
			if node.isLeftChild() {
				node.Parent.Left = node.Right
				node.Right.Parent = node.Parent
			} else {
				node.Parent.Right = node.Left
				node.Right.Parent = node.Parent
			}
		} else {
			root = node.Right
		}
		node = nil
	} else {
		replaceNode := node.Right
		for replaceNode.Left != nil {
			replaceNode = replaceNode.Left
		}
		node.Val = replaceNode.Val
		root, _ = root.delete(replaceNode)
	}
	return root, true
}

func (root *BinaryNode) isLeftChild() bool {
	if root.Parent == nil {
		return false
	}
	return root.Parent.Left == root
}
