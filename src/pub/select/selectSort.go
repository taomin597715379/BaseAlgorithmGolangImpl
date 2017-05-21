package pub

type IntSlice []int

func (s IntSlice) Len() int {
	return len(s)
}
func (s IntSlice) Less(i, j int) bool {
	return s[i] > s[j]
}
func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func selectSort(a IntSlice) IntSlice {
	var i, j int
	var n = len(a)
	if n == 0 || n == 1 {
		return a
	}
	for i = 0; i < n-1; i++ {
		for j = i + 1; j < n; j++ {
			if a.Less(i, j) {
				a.Swap(i, j)
			}
		}
	}
	return a
}
