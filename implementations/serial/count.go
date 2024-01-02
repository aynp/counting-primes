package serial

import (
	"github.com/aynp/counting-primes/lib"
)

func CountPrimes(n int) int {
	count := 0
	for i := 2; i < n; i++ {
		if lib.IsPrime(i) {
			count++
		}
	}
	return count
}
