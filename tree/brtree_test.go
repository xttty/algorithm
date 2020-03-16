package tree_test

import (
	"algorithm/tree"
	"testing"
)

func TestBRTree(t *testing.T) {
	// root := tree.BinaryNode{
	// 	Val: 100,
	// }
	root := tree.BlackRedTreeRoot(100)
	root, _ = root.Add(10)
	root, _ = root.Add(60)
	root, _ = root.Add(160)
	root, _ = root.Add(460)
	root, _ = root.Add(360)
	// root.Add(50)
	// root.Add(21)

	node := root.Search(460)
	t.Log(node)
	// root.Delete(160)
	// node = root.Search(160)
	// t.Log(node)
}
