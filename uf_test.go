package algorithm

import "testing"

func TestUF(t *testing.T) {
	uf := NewUF(10)
	uf.Connection(0, 4)
	uf.Connection(1, 4)
	uf.Connection(3, 5)
	uf.Connection(7, 9)
	uf.Connection(1, 3)
	uf.Connection(4, 3)
	uf.Connection(2, 4)
	t.Log(uf.IsConnection(9, 3))
	t.Log(uf.IsConnection(2, 5))
}
