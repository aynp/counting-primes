package range_cluster_channel

import (
	"github.com/aynp/counting-primes/config"
	"github.com/aynp/counting-primes/lib"
)

func countPrimesInRange(start, end int, ch chan int) {
	count := 0

	for i := start; i < end; i++ {
		if lib.IsPrime(i) {
			count++
		}
	}

	ch <- count
}

func CountPrimes(n int) int {
	thread_count := config.NUM_THREADS

	count := 0
	chunk_size := n / thread_count

	ch := make(chan int, thread_count)

	for i := 0; i < thread_count; i++ {
		go countPrimesInRange(i*chunk_size, (i+1)*chunk_size, ch)
	}

	for i := 0; i < thread_count; i++ {
		count += <-ch
	}

	return count
}
