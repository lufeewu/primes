package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
)

// CalcPrime1 calc prime1
func CalcPrime1(sum int64) (int64, int64) {
	var i int64 = 2
	for ; i <= int64(math.Sqrt(float64(sum))); i++ {
		if sum%i == 0 {
			return i, sum / i
		}
	}
	return 0, 0
}

// PrimeResult prime result
type PrimeResult struct {
	Product int64
	Prime1  int64
	Prime2  int64
}

// CalcPrime2 calc prime2
func CalcPrime2(productMin int64, productMax int64) []PrimeResult {
	var result []PrimeResult
	var s sync.WaitGroup
	var mutex sync.Mutex
	for product := productMin; product <= productMax; product++ {
		s.Add(1)
		go func(product int64) {
			count := 0
			var res int64
			for i := int64(2); i <= int64(math.Sqrt(float64(product))); i = i + 1 {
				if product%i == 0 {
					count++
					res = i
				}
				if count > 1 {
					break
				}
			}
			if count == 1 {
				mutex.Lock()
				result = append(result, PrimeResult{
					Product: product,
					Prime1:  res,
					Prime2:  product / res,
				})
				mutex.Unlock()
			}
			s.Done()
		}(product)
	}
	s.Wait()
	return result
}

func main() {
	fmt.Println(CalcPrime1(7140229933))
	res := CalcPrime2(6541367000, 6541367999)
	// fmt.Println(len(res))
	sort.SliceStable(res, func(i, j int) bool {
		return res[i].Product < res[j].Product
	})
	for i := 0; i < len(res); i++ {
		fmt.Printf("%d = %d * %d\n", res[i].Product, res[i].Prime1, res[i].Prime2)
	}
}
