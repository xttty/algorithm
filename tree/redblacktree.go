package tree

const red = true
const black = false

// RBNode 红黑树节点
type RBNode struct {
	Val    int
	Left   *RBNode
	Right  *RBNode
	Parent *RBNode
	color  bool
}

// RedBlackTreeRoot 创建红黑树祖节点
func RedBlackTreeRoot(value int) *RBNode {
	return &RBNode{
		Val:   value,
		color: black,
	}
}

// Search 红黑树查找
func (root *RBNode) Search(value int) *RBNode {
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

// Add 红黑树的插入
func (root *RBNode) Add(value int) (*RBNode, bool) {
	node := root
	newNode := &RBNode{
		Val:   value,
		color: red,
	}
	for node != nil {
		if node.Val == value {
			return root, false
		} else if node.Val < value {
			if node.Right == nil {
				newNode.Parent = node
				node.Right = newNode
				break
			} else {
				node = node.Right
			}
		} else {
			if node.Left == nil {
				newNode.Parent = node
				node.Left = newNode
				break
			} else {
				node = node.Left
			}
		}
	}
	root = root.fixAfterAdd(newNode)
	return root, true
}

func (root *RBNode) fixAfterAdd(node *RBNode) *RBNode {
	for node != nil && node.Parent != nil && node.Parent.color == red {
		if node.Parent.isLeftChild() { //父节点是左子节点
			if node.isLeftChild() {
				if node.Parent.Parent.Left != nil && node.Parent.Parent.Left.color == red && node.Parent.Parent.Right != nil && node.Parent.Parent.Right.color == red { //父节点和叔节点都是红色
					node.Parent.Parent.Left.color = black
					node.Parent.Parent.Right.color = black
					if node.Parent.Parent != root {
						node.Parent.Parent.color = red
					}
					node = node.Parent.Parent
				} else { //父节点是红色 叔节点是黑色
					root = root.rightRotation(node.Parent.Parent)
					node.Parent.color = black
					node.Parent.Right.color = red
				}
			} else {
				root = root.leftRotation(node.Parent)
				node = node.Left
			}
		} else { //父节点是右子节点
			if !node.isLeftChild() {
				if node.Parent.Parent.Left != nil && node.Parent.Parent.Left.color == red && node.Parent.Parent.Right != nil && node.Parent.Parent.Right.color == red {
					node.Parent.Parent.Left.color = black
					node.Parent.Parent.Right.color = black
					if node.Parent.Parent != root {
						node.Parent.Parent.color = red
					}
					node = node.Parent.Parent
				} else {
					root = root.leftRotation(node.Parent.Parent)
					node.Parent.color = black
					node.Parent.Left.color = red
				}
			} else {
				root = root.rightRotation(node.Parent)
				node = node.Right
			}
		}
	}
	return root
}

// Delete 红黑树删除
func (root *RBNode) Delete(value int) (*RBNode, bool) {
	node := root.Search(value)
	if node == nil {
		return root, false
	}
	return root.delete(node)
}

func (root *RBNode) delete(node *RBNode) (*RBNode, bool) {
	if node.Right != nil && node.Left != nil {
		replaceNode := node.Right
		for replaceNode.Left != nil {
			replaceNode = replaceNode.Right
		}
		node.Val = replaceNode.Val
		node = replaceNode
	}
	root = root.fixBeforeDelete(node)
	if node == root {
		return nil, true
	}
	if node.isLeftChild() {
		node.Parent.Left = nil
	} else {
		node.Parent.Right = nil
	}
	return root, true
}

func (root *RBNode) fixBeforeDelete(node *RBNode) *RBNode {
	for node != nil && node.color == black && node.Parent != nil {
		if node.Right == nil && node.Left != nil {
			node.Val, node.Right.Val = node.Right.Val, node.Val
			node = node.Right
		} else if node.Right != nil && node.Left == nil {
			node.Val, node.Left.Val = node.Left.Val, node.Val
			node = node.Left
		} else {
			if node.isLeftChild() {
				sib := node.Parent.Right
				if sib.color == black {
					if (sib.Right == nil || sib.Left.color == black) && (sib.Left == nil || sib.Left.color == black) {
						if node.Parent.color == red {
							sib.color = red
							break
						} else {
							sib.color = red
							node.Parent.color = black
							node = node.Parent
						}
					} else {
						if sib.Right != nil && sib.Right.color == red {
							root = root.leftRotation(node.Parent)
							node.Parent.color, sib.color = sib.color, node.Parent.color
							sib.Left.color = black
							break
						} else if sib.Left != nil && sib.Left.color == red {
							root = root.rightRotation(sib)
							sib.color, sib.Parent.color = sib.Parent.color, sib.color
						}
					}
				} else {
					root = root.leftRotation(node.Parent)
					sib.color, node.Parent.color = sib.Parent.color, sib.color
				}
			} else {
				sib := node.Parent.Left
				if sib.color == black {
					if (sib.Left == nil || sib.Left.color == black) && (sib.Right == nil || sib.Right.color == black) {
						if node.Parent.color == red {
							sib.color = red
							break
						} else {
							sib.color = red
							node.Parent.color = black
							node = node.Parent
						}
					} else {
						if sib.Left != nil && sib.Left.color == red {
							root = root.rightRotation(node.Parent)
							sib.color, node.Parent.color = node.Parent.color, sib.color
							sib.Left.color = black
							break
						} else if sib.Right != nil && sib.Right.color == red {
							root = root.leftRotation(sib)
							sib.color, sib.Parent.color = sib.Parent.color, sib.color
						}
					}
				} else {
					root = root.rightRotation(node.Parent)
					sib.color, node.Parent.color = node.Parent.color, sib.color
				}
			}
		}
	}
	return root
}

func (root *RBNode) leftRotation(node *RBNode) *RBNode {
	changeRoot := false
	if root == node {
		changeRoot = true
	}
	rightChild := node.Right
	if rightChild == nil {
		return root
	}
	if node.Parent != nil {
		if node.isLeftChild() {
			node.Parent.Left = rightChild
		} else {
			node.Parent.Right = rightChild
		}
	}
	rightChild.Parent = node.Parent
	node.Parent = rightChild
	node.Right = rightChild.Left
	if rightChild.Left != nil {
		rightChild.Left.Parent = node
	}
	rightChild.Left = node
	if changeRoot {
		root = rightChild
	}
	return root
}

func (root *RBNode) rightRotation(node *RBNode) *RBNode {
	changeRoot := false
	if root == node {
		changeRoot = true
	}
	leftChild := node.Left
	if leftChild == nil {
		return root
	}
	if node.Parent != nil {
		if node.isLeftChild() {
			node.Parent.Left = leftChild
		} else {
			node.Parent.Right = leftChild
		}
	}
	leftChild.Parent = node.Parent
	node.Parent = leftChild
	node.Left = leftChild.Right
	if leftChild.Right != nil {
		leftChild.Right.Parent = node
	}
	leftChild.Right = node
	if changeRoot {
		root = leftChild
	}
	return root
}

func (root *RBNode) isLeftChild() bool {
	if root.Parent == nil {
		return false
	}
	return root.Parent.Left == root
}
