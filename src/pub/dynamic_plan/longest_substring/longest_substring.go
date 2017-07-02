package pub

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 0 {
		return 0
	}
	var start, ret int
	var hashmap = make(map[byte]int, 0)
	for k, v := range []byte(s) {
		if _, ok := hashmap[v]; ok && hashmap[v] >= start {
			start = hashmap[v] + 1
			delete(hashmap, v)
		}
		hashmap[v] = k
		if ret < (k - start + 1) {
			ret = k - start + 1
		}
	}
	return ret
}
