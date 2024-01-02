package range_cluster_global_mutex

import (
	"sync"
	"sync/atomic"

	"github.com/aynp/counting-primes/config"
	"github.com/aynp/counting-primes/lib"
)

var wg sync.WaitGroup
var count atomic.Int64

func countPrimesInRange(start, end int) {
	defer wg.Done()

	for i := start; i < end; i++ {
		if lib.IsPrime(i) {
			count.Add(1)
		}
	}
}

func CountPrimes(n int) int {
	count.Store(0)

	thread_count := config.NUM_THREADS
	chunk_size := n / thread_count

	for i := 0; i < thread_count; i++ {
		wg.Add(1)
		go countPrimesInRange(i*chunk_size, (i+1)*chunk_size)
	}

	wg.Wait()

	return int(count.Load())
}
