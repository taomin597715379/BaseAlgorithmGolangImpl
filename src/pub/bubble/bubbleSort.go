package pub

func bubbleSort(a []int) []int {
	var i, j int
	var n = len(a)
	if n == 0 || n == 1 {
		return a
	}
	for i = 0; i < n-1; i++ {
		for j = 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

func bubbleSortImprv(a []int) []int {
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
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}
	}
	return a
}
