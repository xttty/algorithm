package tree_test

import (
	"algorithm/tree"
	"testing"
)

func TestBRTree(t *testing.T) {
	// root := tree.BinaryNode{
	// 	Val: 100,
	// }
	// root.Add(50)
	// root.Add(80)
	// root.Add(23)
	// root.Add(110)
	// root.Add(180)
	// root.Add(200)
	// root.Delete(110)

	root := tree.RedBlackTreeRoot(100)
	node := root
	root, _ = root.Add(10)
	root, _ = root.Add(60)
	root, _ = root.Add(160)
	root, _ = root.Add(460)
	root, _ = root.Add(360)
	root, _ = root.Add(500)
	root, _ = root.Add(450)
	// root, _ = root.Add(111)

	root, _ = root.Delete(60)

	node = root.Search(500)
	t.Log(node)
	// root.Delete(160)
	// node = root.Search(160)
	// t.Log(node)
}
