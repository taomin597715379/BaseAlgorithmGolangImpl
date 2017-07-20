package pub

import (
	"fmt"
	"testing"
)

func Test_longComSubStr_1(t *testing.T) {
	fmt.Println(longComSubStr(`abcdsacc`, `bdsccd`))
}

func Test_longComSubStr_2(t *testing.T) {
	fmt.Println(longComSubStr(`abc`, `a`))
}

func Test_longComSubStr_3(t *testing.T) {
	fmt.Println(longComSubStr(`abc`, ``))
}

func Test_longComSubStr_4(t *testing.T) {
	fmt.Println(longComSubStr(``, ``))
}

func Test_longComSubStrAndOutPut_1(t *testing.T) {
	longComSubStrAndOutPut(`abcdefg`, `xyzabcd`)
}

func Test_longComSubStrAndOutPut_2(t *testing.T) {
	longComSubStrAndOutPut(`abcdsacc`, `bdsccd`)
}

func Test_longComSubSeqAndOutput(t *testing.T) {
	longComSubSeqAndOutput(`abcdsacc`, `bdsccd`)
}
