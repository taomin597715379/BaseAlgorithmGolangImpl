// LSUB(x,y) = LSUB(x-1,y-1)+1; 若S1[x]==S2[y]
// LSUB(x,y) = 0; 若S1[x]!=S2[y]
// S1: abcdefg
// S2: xyzabcd

package pub

const MAX int = 10000

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func longComSubStr(str1, str2 string) int {
	var l [MAX][MAX]int
	var result int
	var str1toByte []byte = []byte(str1)
	var str2toByte []byte = []byte(str2)
	for i := 0; i <= len(str1toByte); i++ {
		for j := 0; j <= len(str2toByte); j++ {
			if i == 0 || j == 0 {
				l[i][j] = 0
			} else if str1toByte[i-1] == str2toByte[j-1] {
				l[i][j] = l[i-1][j-1] + 1
				result = max(l[i][j], result)
			} else {
				l[i][j] = 0
			}
		}
	}
	return result
}
