package algorithm

// UF 并查集结构
type UF struct {
	collection []int
}

// NewUF 新建并查集
func NewUF(cnt int) *UF {
	uf := &UF{
		collection: make([]int, cnt),
	}
	for i := 0; i < cnt; i++ {
		uf.collection[i] = i
	}
	return uf
}

// Connection i，j联通
func (uf *UF) Connection(i, j int) {
	rootI, deepI := uf.root(i)
	rootJ, deepJ := uf.root(j)
	if deepI > deepJ {
		uf.collection[j] = rootI
	} else {
		uf.collection[i] = rootJ
	}
}

// IsConnection 判断i，j是否联通
func (uf *UF) IsConnection(i, j int) bool {
	rootI, _ := uf.root(i)
	rootJ, _ := uf.root(j)
	return rootI == rootJ
}

func (uf *UF) root(i int) (int, int) {
	root, deep := i, 1
	for root != uf.collection[root] {
		temp := root
		root = uf.collection[root]
		uf.collection[temp] = uf.collection[root]
		deep++
	}
	return root, deep
}
