package hyparview

import (
	cr "crypto/rand"
	"math/big"
	"math/rand"
)

type RandomSource interface {
	Intn(n int) int
}

type defaultRandom struct{}

func (d defaultRandom) Intn(n int) int {
	return rand.Intn(n)
}

var DefaultRandom RandomSource = defaultRandom{}

// rand [0, n] inclusive
func Rint64Crypto(n int64) int64 {
	bn := new(big.Int).SetInt64(int64(n + 1))
	bi, _ := cr.Int(cr.Reader, bn)
	return bi.Int64()
}

// rand [0, n] inclusive
func RintCrypto(n int) int {
	return int(Rint64Crypto(int64(n)))
}

// RintWithSource uses the provided RandomSource, falls back to DefaultRandom if nil
// rand [0, n] inclusive
func RintWithSource(n int, src RandomSource) int {
	if src == nil {
		src = DefaultRandom
	}
	return src.Intn(n + 1)
}

// rint is a placeholder so we can swap out for rintCrypto in testing
// rand [0, n] inclusive
func Rint(n int) int {
	return RintWithSource(n, nil)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
