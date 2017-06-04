package pub

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 动态规划
func longIncrSubseq(a []int) int {
	var i, j int
	var res int = 1
	var n int = len(a)
	var dp = make([]int, n)
	for i = 0; i < n; i++ {
		dp[i] = 0
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	for i = 1; i < n; i++ {
		for j = 0; j < i; j++ {
			if a[j] < a[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

// 分治
