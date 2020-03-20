package algorithm_test

import (
	"algorithm"
	"testing"
)

func TestKMP(t *testing.T) {
	str := "ababaffabaababababbabaabababababababaaaasdfbabababaaaasdf"
	needle := "abababababababaaaasdf"
	pos, ok := algorithm.KMPSearch(needle, str)
	t.Log(pos, ok, str[pos:pos+len(needle)])
}
