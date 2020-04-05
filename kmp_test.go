package algorithm

import (
	"testing"
)

func TestKMP(t *testing.T) {
	str := "ababaffabaababababbabaabababababababaaaasdfbabababaaaasdf"
	needle := "abababababababaaaasdf"
	pos, ok := KMPSearch(needle, str)
	t.Log(pos, ok, str[pos:pos+len(needle)])
}
