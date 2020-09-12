package math

import "math"

// IsPrimeNaive uses the naive method for determining if the number is prime.
func IsPrimeNaive(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func IsPrimeSieveOfEratosthenes(n int) bool {
	// Set all values to true other than 0 and 1
	hash := make([]bool, n + 1)
	for i := 2; i <= n; i++ {
		hash[i] = true
	}

	// Cross off remaining multiples of `prime`
	crossOff := func(prime int) {
		for i := prime * prime; i < len(hash); i += prime {
			hash[i] = false
		}
	}

	// Get the next prime number in the hash
	nextPrime := func(prime int) int {
		next := prime + 1
		for next < len(hash) && !hash[next] {
			next++
		}
		return next
	}

	// Record all the primes in the hash
	prime := 2
	for prime <= int(math.Sqrt(float64(n))) {
		crossOff(prime)
		prime = nextPrime(prime)
	}

	// Return whether n is prime
	return hash[len(hash) - 1]
}
