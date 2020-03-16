package tree

const red = true
const black = false

// BRNode 红黑树节点
type BRNode struct {
	Val    int
	Left   *BRNode
	Right  *BRNode
	Parent *BRNode
	color  bool
}

// BlackRedTreeRoot 创建红黑树祖节点
func BlackRedTreeRoot(value int) *BRNode {
	return &BRNode{
		Val:   value,
		color: black,
	}
}

// Search 红黑树查找
func (root *BRNode) Search(value int) *BRNode {
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
func (root *BRNode) Add(value int) (*BRNode, bool) {
	node := root
	newNode := &BRNode{
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

func (root *BRNode) fixAfterAdd(node *BRNode) *BRNode {
	for node != nil && node.Parent != nil && node.Parent.color == red {
		if node.Parent.isLeftChild() { //父节点是左子节点
			if node.Parent.Parent.Left != nil && node.Parent.Parent.Left.color == red && node.Parent.Parent.Right != nil && node.Parent.Parent.Right.color == red { //父节点和叔节点都是红色
				if node.Parent.Parent.Left != nil {
					node.Parent.Parent.Left.color = black
				}
				if node.Parent.Parent.Right != nil {
					node.Parent.Parent.Right.color = black
				}
				if node.Parent.Parent != root {
					node.Parent.Parent.color = red
				}
				node = node.Parent.Parent
			} else { //父节点是红色 叔节点是黑色
				if !node.isLeftChild() {
					root = root.leftRotation(node.Parent)
					node = node.Left
				} else {
					root = root.rightRotation(node.Parent.Parent)
					node.Parent.color = black
					node.Parent.Right.color = red
				}
			}
		} else { //父节点是右子节点
			if node.Parent.Parent.Left != nil && node.Parent.Parent.Left.color == red && node.Parent.Parent.Right != nil && node.Parent.Parent.Right.color == red {
				if node.Parent.Parent.Left != nil {
					node.Parent.Parent.Left.color = black
				}
				if node.Parent.Parent.Right != nil {
					node.Parent.Parent.Right.color = black
				}
				if node.Parent.Parent != root {
					node.Parent.Parent.color = red
				}
				node = node.Parent.Parent
			} else {
				if node.isLeftChild() {
					root = root.rightRotation(node.Parent)
					node = node.Right
				} else {
					root = root.leftRotation(node.Parent.Parent)
					node.Parent.color = black
					node.Parent.Left.color = red
				}
			}
		}
	}
	return root
}

func (root *BRNode) leftRotation(node *BRNode) *BRNode {
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
	rightChild.Left = node
	if changeRoot {
		root = rightChild
	}
	return root
}

func (root *BRNode) rightRotation(node *BRNode) *BRNode {
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
	leftChild.Right = node
	if changeRoot {
		root = leftChild
	}
	return root
}

func (root *BRNode) isLeftChild() bool {
	if root.Parent == nil {
		return false
	}
	return root.Parent.Left == root
}
