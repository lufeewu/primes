package main

import (
	"testing"
)

func BenchmarkPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := CalcPrime2(6541367000, 6541367999)
		if len(res) != 151 {
			b.Errorf("error, len=%d", len(res))
		}
	}

}
