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

func bubbleSort(a IntSlice) IntSlice {
	var i, j int
	var n = len(a)
	if n == 0 || n == 1 {
		return a
	}
	for i = 0; i < n-1; i++ {
		for j = 0; j < n-i-1; j++ {
			if a.Less(j, j+1) {
				a.Swap(j, j+1)
			}
		}
	}
	return a
}

func bubbleSortImprv(a IntSlice) IntSlice {
	var i, j int
	var flag bool = true
	var n = len(a)
	if n == 0 || n == 1 {
		return a
	}
	for i = 0; i < n-1; i++ {
		if flag == false {
			return a
		}
		flag = false
		for j = 0; j < n-i-1; j++ {
			if a.Less(j, j+1) {
				a.Swap(j, j+1)
				flag = true
			}
		}
	}
	return a
}
