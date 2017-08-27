package binary

// 10个写二分查找的9个错。。。。
// 1.求任意一个i，满足a[i] = key，若不存在，则返回-1
func binary_search_1(a []int, n, key int) int {
	var m, i int
	var j int = n - 1
	for i <= j {
		m = i + (j-i)/2
		if a[m] < key {
			i = m + 1
		}
		if a[m] > key {
			j = m - 1
		}
		if a[m] == key {
			return m
		}
	}
	return -1
}

// 1.求最小的i，使得a[i] = key，若不存在，则返回-1
func binary_search_2(a []int, n, key int) int {
	var m, i int
	var j int = n - 1
	for i < j {
		m = i + (j-i)/2
		if a[m] < key {
			i = m + 1
		} else {
			j = m
		}
	}
	if a[j] == key {
		return j
	}
	return -1
}

// 求最大的i，使得a[i] = key，若不存在，则返回-1
func binary_search_3(a []int, n, key int) int {
	var m, i int
	var j int = n - 1
	for i < j {
		m = i + (j-i+1)/2
		if a[m] <= key {
			i = m
		} else {
			j = m - 1
		}
	}
	if a[i] == key {
		return i
	}
	return -1
}

// 求最小的i，使得a[i] > key，若不存在，则返回-1
func binary_search_4(a []int, n, key int) int {
	var m, i int
	var j int = n - 1
	for i < j {
		m = i + (j-i)/2
		if a[m] <= key {
			i = m + 1
		} else {
			j = m
		}
	}
	if a[j] > key {
		return j
	}
	return -1
}

// 求最大的i，使得a[i] < key，若不存在，则返回-1
func binary_search_5(a []int, n, key int) int {
	var m, i int
	var j int = n - 1
	for i < j {
		m = i + (j-i+1)/2
		if a[m] < key {
			i = m
		} else {
			j = m - 1
		}
	}
	if a[i] < key {
		return i
	}
	return -1
}
