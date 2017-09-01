/*
Bloom Filter:

Bloom Filter is a space-efficient random data structure that uses a
bit array to represent a collection succinctly and to determine whether
an element belongs to that collection. This efficiency of Bloom Filter is
a price: in determining whether an element belongs to a collection, it is
possible to mistaken the element that does not belong to this collection to
belong to this set (false positive). Therefore, Bloom Filter is not suitable
for those "zero error" applications. And in applications that tolerate low
error rates, Bloom Filter makes a significant savings in storage space
with minimal errors.

Bloom Filter structure:

        +- x  -+                          +- y -+
       /        \                       /        \
      +-----+-----+-----+----+-----+----+----+-----+-----+-----+-----+-----+----+----+
      |  0  |  0  |  0  |  0 |  0  |  0 | 0  |  0  |  0  |  0  |   0 |  0  |  0 |  0 |
      +-----+-----+-----+----+-----+----+----+-----+-----+-----+-----+-----+----+----+

*/

package bloom

import (
	"fmt"
	"github.com/spaolacci/murmur3"
)

/*
BloomFilter esti_n descript the estimate of the amount of data inserted and
m descript the number of BitSet. Calculate the k value according to the formula:
    k = ln(2)* m/n
the probability of an error when k = ln (2) * m / n is minimal and reference
document: http://pages.cs.wisc.edu/~cao/papers/summary-cache/node8.html.
*/
type BloomFilter struct {
	m uint
	k uint
	b *BitSet
}

type BitSet struct {
	length uint
	set    []uint64
}

func max(x, y uint) uint {
	if x > y {
		return x
	}
	return y
}

func numberBitWords(l uint) int {
	if uint64(l) > ((^uint64(0)) - 64 + 1) {
		return int((^uint64(0)) >> 6)
	}
	return int((l + 64 - 1) >> 6)
}

func newBitSet(l uint) *BitSet {
	nums := numberBitWords(l)
	return &BitSet{
		l,
		make([]uint64, nums),
	}
}

func NewBloom(m, k uint) *BloomFilter {
	return &BloomFilter{
		max(1, m),
		max(1, k),
		newBitSet(m),
	}
}

func (bf *BloomFilter) Set(i uint) {
	bf.b = bf.b.Set(i)
}

func (bf *BloomFilter) Clear(i uint) {
	bf.b = bf.b.Clear(i)
}

func (bf *BloomFilter) Add(data []byte) *BloomFilter {
	h := hashF(data)
	for i := uint(0); i < bf.k; i++ {
		bf.b.Set(uint(location(h, i) % uint64(bf.m)))
	}
	return bf
}

func (bf *BloomFilter) AddString(data string) *BloomFilter {
	return bf.Add([]byte(data))
}

func (b *BitSet) mayBeNeedMoreBit(i uint) {
	numOf64 := numberBitWords(i + 1)
	if b.set == nil {
		b.set = make([]uint64, numOf64)
	} else if cap(b.set) >= numOf64 {
		b.set = b.set[:numOf64]
	} else if cap(b.set) < numOf64 {
		newBitSet := make([]uint64, numOf64, 2*numOf64)
		copy(newBitSet, b.set)
		b.set = newBitSet
	}
	b.length = i + 1
}

func (b *BitSet) Set(i uint) *BitSet {
	if i >= b.length {
		b.mayBeNeedMoreBit(i)
	}
	b.set[i>>6] = b.set[i>>6] | (1 << (i & (64 - 1)))
	return b
}

func (b *BitSet) Clear(i uint) *BitSet {
	if i >= b.length {
		return b
	}
	b.set[i>>6] &^= 1 << (i & (64 - 1))
	return b
}

func hashF(data []byte) [4]uint64 {
	a1 := []byte{1}
	hasher := murmur3.New128()
	hasher.Write(data)
	v1, v2 := hasher.Sum128()
	hasher.Write(a1)
	v3, v4 := hasher.Sum128()
	return [4]uint64{
		v1, v2, v3, v4,
	}
}

func location(h [4]uint64, i uint) uint64 {
	ii := uint64(i)
	return h[ii%2] + ii*h[2+(((ii+(ii%2))%4)/2)]
}
