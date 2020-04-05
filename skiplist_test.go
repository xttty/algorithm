package algorithm

import (
	"testing"
)

func TestSkipList(t *testing.T) {
	head := NewSkipList(100)
	head = head.Add(15)
	head = head.Add(151)
	head = head.Add(34)
	head = head.Add(99)
	head = head.Add(200)
	head = head.Add(460)
	head = head.Delete(100)
	node := head.Search(100)
	t.Log(node)
}
