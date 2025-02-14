package calc

// BinomialCoefficient erwartet zwei Zahlen n und k und liefert den
// Binomialkoeffizienten "n über k".
func BinomialCoefficient(n, k int) int {
	// TODO
	return Factorial(n) / (Factorial(n-k) * Factorial(k))
}

// calculates factorial of n
func Factorial(n int) int {

	if n <= 0 {

		return 1

	}

	return n * Factorial(n-1)

}

// n!/(n-k)!*k!
