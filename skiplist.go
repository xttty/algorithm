package algorithm

import (
	"math/rand"
	"time"
)

const maxLevel = 7

// SkipNode  跳跃表节点
type SkipNode struct {
	Val   int
	level uint8
	ptrs  []*SkipNode
}

// NewSkipList 初始化
func NewSkipList(value int) *SkipNode {
	level := randomLevel()
	node := &SkipNode{
		Val:   value,
		level: level,
		ptrs:  make([]*SkipNode, level),
	}
	return node
}

// Search 跳跃表搜索
func (head *SkipNode) Search(value int) *SkipNode {
	node := head.search(value)
	if node.Val != value {
		return nil
	}
	return node
}

func (head *SkipNode) search(value int) *SkipNode {
	node := head
	level := node.level
	for node.Val < value {
		for node.ptrs[level-1] == nil || node.ptrs[level-1].Val > value {
			level--
			if level <= 0 {
				break
			}
		}
		if level == 0 || node.ptrs[level-1] == nil {
			break
		}
		node = node.ptrs[level-1]
		level = node.level
	}
	return node
}

// Add 新增节点
func (head *SkipNode) Add(value int) *SkipNode {
	node := head.search(value)
	if node.Val == value {
		return head
	} else if node.Val > value {
		node.Val, value = value, node.Val
	}
	insertNode := &SkipNode{
		Val:   value,
		level: randomLevel(),
	}
	insertNode.ptrs = make([]*SkipNode, insertNode.level)
	node = head
	for node.ptrs[0] != insertNode {
		level := node.level
		for level > insertNode.level {
			level--
		}
		for level > 0 && (node.ptrs[level-1] == nil || node.ptrs[level-1].Val > value) {
			insertNode.ptrs[level-1] = node.ptrs[level-1]
			node.ptrs[level-1] = insertNode
			level--
		}
		if level > 0 {
			node = node.ptrs[level-1]
		}
	}
	return head
}

// Delete 删除节点
func (head *SkipNode) Delete(value int) *SkipNode {
	node := head.search(value)
	if node.Val != value {
		return head
	}
	tempNode := head
	level := uint8(maxLevel)
	for level > 0 && tempNode != nil {
		level = tempNode.level
		for level > node.level {
			level--
		}
		for level > 0 && tempNode.ptrs[level-1] == node {
			tempNode.ptrs[level-1] = node.ptrs[level-1]
			level--
		}
		if level > 0 {
			tempNode = tempNode.ptrs[level-1]
		}
	}
	return head
}

func randomLevel() uint8 {
	rand.Seed(time.Now().UnixNano())
	level := rand.Int()%maxLevel + 1
	return uint8(level)
}
