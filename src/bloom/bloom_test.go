package bloom

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

// 3 < 10
func Test_Bloom_1(t *testing.T) {
	bf := NewBloom(10, 2)
	bf.Set(3)
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, bf.b.set)
	fmt.Println(buf.Bytes())
	fmt.Println(len(bf.b.set))
	fmt.Println(bf.b.set)
}

// 3 < 10 and clear()
func Test_Bloom_2(t *testing.T) {
	bf := NewBloom(10, 2)
	bf.Set(3)
	bf.Clear(3)
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, bf.b.set)
	fmt.Println(buf.Bytes())
	fmt.Println(len(bf.b.set))
	fmt.Println(bf.b.set)
}

// 12 > 10
func Test_Bloom_3(t *testing.T) {
	bf := NewBloom(10, 2)
	bf.Set(12)
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, bf.b.set)
	fmt.Println(buf.Bytes())
	fmt.Println(len(bf.b.set))
	fmt.Println(bf.b.set)
}

// 65 > 64(word) && 65 > 10
func Test_Bloom_4(t *testing.T) {
	bf := NewBloom(10, 2)
	bf.Set(65)
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, bf.b.set)
	fmt.Println(buf.Bytes())
	fmt.Println(len(bf.b.set))
	fmt.Println(bf.b.set)
}

// addstring test
func Test_Bloom_5(t *testing.T) {
	bf := NewBloom(10, 2)
	bf.AddString(`hello world`)
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, bf.b.set)
	fmt.Println(buf.Bytes())
	fmt.Println(len(bf.b.set))
	fmt.Println(bf.b.set)
}
