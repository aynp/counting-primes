package range_cluster_channel

import (
	"sync"
	"sync/atomic"

	"github.com/aynp/counting-primes/config"
	"github.com/aynp/counting-primes/lib"
)

var current atomic.Int32
var totalPrimes atomic.Int32

func worker(max_test int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		i := current.Add(1)
		if i > int32(max_test) {
			break
		}

		if lib.IsPrime(int(i)) {
			totalPrimes.Add(1)
		}
	}
}

func CountPrimes(max_test int) int {
	thread_count := config.NUM_THREADS

	var wg sync.WaitGroup

	current.Store(0)
	totalPrimes.Store(0)

	for i := 0; i < thread_count; i++ {
		go worker(max_test, &wg)
	}

	return int(current.Load())
}
