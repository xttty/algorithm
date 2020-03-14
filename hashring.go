package algorithm

import (
	"hash"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

type node struct {
	nodeKey string
	nodeV   uint32
}

type nodes []*node

func (a nodes) Len() int           { return len(a) }
func (a nodes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a nodes) Less(i, j int) bool { return a[i].nodeV < a[j].nodeV }

var defaultVirtualNodes int = 1000

// HashRing 一致性hash环
type HashRing struct {
	virtualNodes int
	weight       map[string]int
	nodesList    nodes
	mu           sync.Mutex
}

// NewHashRing new方法
func NewHashRing(virtual int) *HashRing {
	if virtual == 0 {
		virtual = defaultVirtualNodes
	}
	hr := &HashRing{
		virtualNodes: virtual,
		weight:       make(map[string]int),
	}
	hr.genarate()
	return hr
}

// Update 更新
func (hr *HashRing) Update(name string, weight int) {
	hr.mu.Lock()
	defer hr.mu.Unlock()

	hr.weight[name] = weight
	hr.genarate()
}

// Delete 删除节点
func (hr *HashRing) Delete(name string) {
	hr.mu.Lock()
	defer hr.mu.Unlock()

	delete(hr.weight, name)
	hr.genarate()
}

// Add 添加节点
func (hr *HashRing) Add(name string, weight int) {
	hr.Update(name, weight)
}

// GetNode 查找对应的节点key
func (hr *HashRing) GetNode(key string) string {
	hr.mu.Lock()
	defer hr.mu.Unlock()

	hashValue := hashValue(key)
	n := sort.Search(hr.nodesList.Len(), func(i int) bool {
		return hashValue <= hr.nodesList[i].nodeV
	})
	return hr.nodesList[n].nodeKey
}

func (hr *HashRing) genarate() {
	allWeight := 0
	for _, w := range hr.weight {
		allWeight += w
	}
	for key, w := range hr.weight {
		var allocNodes int = int(float32(w) / float32(allWeight) * float32(hr.virtualNodes))
		for i := 0; i < allocNodes; i++ {
			vNode := &node{
				nodeKey: key,
				nodeV:   hashValue(key + ":" + strconv.Itoa(i)),
			}
			hr.nodesList = append(hr.nodesList, vNode)
		}
	}
	sort.Sort(hr.nodesList)
}

var hs hash.Hash32 = crc32.NewIEEE()

func hashValue(key string) uint32 {
	hs.Write([]byte(key))
	value := hs.Sum32()
	hs.Reset()
	return value
}
