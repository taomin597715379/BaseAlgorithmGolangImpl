package pub

import (
	"fmt"
	"testing"
)

func Test_lengthOfLongestSubstring_1(t *testing.T) {
	var a string = `abcabcbb`
	fmt.Println(lengthOfLongestSubstring(a))
}

func Test_lengthOfLongestSubstring_2(t *testing.T) {
	var a string = `bbbbbbb`
	fmt.Println(lengthOfLongestSubstring(a))
}

func Test_lengthOfLongestSubstring_3(t *testing.T) {
	var a string = `pwwkew`
	fmt.Println(lengthOfLongestSubstring(a))
}

func Test_lengthOfLongestSubstring_4(t *testing.T) {
	var a string = ``
	fmt.Println(lengthOfLongestSubstring(a))
}
