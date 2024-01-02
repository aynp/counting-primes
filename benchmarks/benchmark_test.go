package benchmarks

import (
	"testing"

	"github.com/aynp/counting-primes/config"
	"github.com/aynp/counting-primes/implementations/precompute"
	"github.com/aynp/counting-primes/implementations/range_cluster_channel"
	"github.com/aynp/counting-primes/implementations/range_cluster_global_mutex"
	"github.com/aynp/counting-primes/implementations/serial"
)

func BenchmarkSerial(b *testing.B) {
	for n := 0; n < b.N; n++ {
		serial.CountPrimes(config.MAX_TEST)
	}
}

func BenchmarkRangeClusterGlobalMutex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		range_cluster_global_mutex.CountPrimes(config.MAX_TEST)
	}
}

func BenchmarkRangeClusterChannel(b *testing.B) {
	for n := 0; n < b.N; n++ {
		range_cluster_channel.CountPrimes(config.MAX_TEST)
	}
}

func BenchmarkFairGlobalMutex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		range_cluster_global_mutex.CountPrimes(config.MAX_TEST)
	}
}

func BenchmarkPrecompute(b *testing.B) {
	for n := 0; n < b.N; n++ {
		precompute.CountPrimes(config.MAX_TEST)
	}
}
