package lib

// IsPrime returns true if n is prime, false otherwise.
// This takes O(sqrt(n)) time.
func IsPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
