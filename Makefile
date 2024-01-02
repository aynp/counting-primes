VERSION ?= 0.1.0

.PHONY: benchmark
benchmark:
	cd benchmarks; go test -bench .
