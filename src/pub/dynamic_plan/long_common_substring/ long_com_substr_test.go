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
