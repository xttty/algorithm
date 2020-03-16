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
func (root *BinaryNode) Delete(value int) bool {
	node := root.Search(value)
	if node == nil {
		return false
	}
	if node.isLeftChild() {
		node.Parent.Left = nil
	} else {
		node.Parent.Right = nil
	}
	return true
}

func (root *BinaryNode) isLeftChild() bool {
	if root.Parent == nil {
		return false
	}
	return root.Parent.Left == root
}
