package pub

func insertSort(a []int) []int {
	var i, j int
	var n = len(a)
	if n == 0 || n == 1 {
		return a
	}
	for i = 1; i < n; i++ {
		for j = i - 1; j >= 0 && a[j] > a[j+1]; j-- {
			a[j], a[j+1] = a[j+1], a[j]
		}
	}
	return a
}
