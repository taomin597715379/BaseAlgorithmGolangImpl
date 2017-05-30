package pub

type IntSlice []int

func (s IntSlice) Len() int {
	return len(s)
}
func (s IntSlice) Less(i, j int) bool {
	return s[i] > s[j]
}

func (s IntSlice) Equal(i, j int) bool {
	return s[i] == s[j]
}
func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func partition(a IntSlice, left int, right int) int {
	var pre_index int = left
	var i, j int = left, right
	for {
		for a.Less(j, pre_index) && j > i {
			j--
		}
		for (a.Less(pre_index, i) || a.Equal(pre_index, i)) && i < j {
			i++
		}
		if i >= j {
			break
		}
		a.Swap(i, j)
	}
	a.Swap(i, pre_index)
	return i
}

func quickSort(a IntSlice, left, right int) {
	if left < right {
		m := partition(a, left, right)
		quickSort(a, left, m-1)
		quickSort(a, m+1, right)
	}
}
