package pub

import (
	"fmt"
)

const MAX int = 10000

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// LSUB(x,y) = LSUB(x-1,y-1)+1; 若S1[x]==S2[y]
// LSUB(x,y) = 0; 若S1[x]!=S2[y]
// S1: abcdefg
// S2: xyzabcd
//最长公共子串
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

//最长公共子串，同时输出结果
func longComSubStrAndOutPut(str1, str2 string) {
	var l [MAX][MAX]int
	var result, maxindex, k int
	var str1toByte []byte = []byte(str1)
	var str2toByte []byte = []byte(str2)
	var strret []byte
	for i := 0; i <= len(str1toByte); i++ {
		for j := 0; j <= len(str2toByte); j++ {
			if i == 0 || j == 0 {
				l[i][j] = 0
			} else if str1toByte[i-1] == str2toByte[j-1] {
				l[i][j] = l[i-1][j-1] + 1
				// result = max(l[i][j], result)
				if result < l[i][j] {
					result = l[i][j]
					maxindex = i - result
				}
			} else {
				l[i][j] = 0
			}
		}
	}
	if result == 0 {
		fmt.Println(``)
	} else {
		k = maxindex
		for result > 0 {
			strret = append(strret, str1toByte[k])
			k++
			result--
		}
		fmt.Println(string(strret))
	}
}

// LSUB(x,y) = LSUB(x-1,y-1)+1; 若S1[x]==S2[y]
// LSUB(x,y) = max{LSUB(x,y-1), LSUB(x-1,y)}; 若S1[x]!=S2[y]
// S1: abcdefg
// S2: xyzabcd
// 最长公共子序列，同时输出序列
func longComSubSeqAndOutput(str1, str2 string) {
	var l [MAX][MAX]int
	var strret [MAX]byte = [MAX]byte{}
	var result, i, j int
	for i = 0; i <= len(str1); i++ {
		for j = 0; j <= len(str2); j++ {
			if i == 0 || j == 0 {
				l[i][j] = 0
			} else if str1[i-1] == str2[j-1] {
				l[i][j] = l[i-1][j-1] + 1
				result = max(l[i][j], result)
			} else {
				l[i][j] = max(l[i][j-1], l[i-1][j])
			}
		}
	}
	i = len(str1)
	j = len(str2)
	var g, k int = l[i][j], l[i][j]
	for i > 0 && j > 0 {
		if str1[i-1] == str2[j-1] && l[i][j] == l[i-1][j-1]+1 {
			k--
			strret[k] = str1[i-1]
			i--
			j--
		} else if str1[i-1] != str2[j-1] && l[i-1][j] > l[i][j-1] {
			i--
		} else {
			j--
		}
	}
	fmt.Println(string(strret[:g]))
}
