# Counting Primes

Comparisons of various multithreaded implimentations of counting prime numbers upto a given numbers.

# Implimentations
- Serial  
  In a single thread, count the number of primes serially.

- Range Cluster Global Mutex  
  For each thread, divide the range of numbers to be checked into equal sized chunks. Each thread will count the number of primes in the chunk assigned to it. After counting the numbers in the chunk update the global count variable.

- Range Cluster Channel
  Same as Range Cluster Global Mutex, but instead of using a global mutex, go channels are used.

- Fair Global Mutex
  Each worker thread will check the what's the next number to be checked and check if it's prime.

- Precompute
  Kind of a cheat. Precomputes the number of primes upto a given number using the go `init` function. Return the precomputed value without any computation.

