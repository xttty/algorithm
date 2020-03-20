package algorithm

// KMPSearch kmp算法实现字符串匹配
// KMP 算法在大学时便看过很多次，工作后又温习过一次，但是总是容易忘记
// 这次试用golang再实现一次，希望加深记忆
// KMP算法核心在于分析子串自身的特性，使得目标串的指针不需要后退，让匹配时间复杂度在O(n)
func KMPSearch(needle, target string) (int, bool) {
	next := kmpNextValues(needle)
	i, j, nLen, tLen := 0, 0, len(needle), len(target)
	ok := false
	pos := 0
	for i < nLen && j < tLen {
		if i == -1 || needle[i] == target[j] {
			i++
			j++
		} else {
			i = next[i]
		}
	}
	if i >= nLen {
		ok = true
		pos = j - nLen
	}
	return pos, ok
}

// kmp nextval 第一个值永远是-1
//  ababc str
//  00120 最长公共前缀  对于str[i] str[0:k-1] == str[i-k-1:i]
// -10012 nextVal “最长公共前缀”数组整体向右移，然后首位填-1
// -10002 优化后的nextVal 可见 next[3] = 1 str[1] == str[3] == 'b' 所以直接next[3] = next[1] = 0
func kmpNextValues(str string) []int {
	k, j, sLen := -1, 0, len(str)
	nextVal := make([]int, sLen)
	nextVal[0] = -1
	for j < sLen-1 {
		if k == -1 || str[j] == str[k] {
			k++
			j++
			if str[j] != str[k] {
				nextVal[j] = k
			} else {
				// 因为当str[j]不匹配时，且str[j] == str[nextval[j]]，则必然也无法匹配，额外增加回溯次数
				nextVal[j] = nextVal[k]
			}
		} else {
			k = nextVal[k]
		}
	}
	return nextVal
}
