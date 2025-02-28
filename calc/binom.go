package calc

// BinomialCoefficient erwartet zwei Zahlen n und k und liefert den
// Binomialkoeffizienten "n über k".
func BinomialCoefficient(n, k int) int {

	// rekursionsanker, wenn
	if k == 0 || n == 0 || k == n {

		return 1

	} else if k == 1 {

		return n
	}

	// ziehen ohne zurücklegen

	// wie viele möglichkeiten ein element aus n verbleibenden elementen zu ziehen?
	// n möglichkeiten

	// wenn gezogen wurde verringert sich die anzahl an versuchen und die anzahl an verbleibenden elementen
	// pro zug werden die möglichkeiten aufaddiert

	// weiter ziehen
	return n + BinomialCoefficient(n-1, k-1)

}

// calculates factorial of n
func Factorial(n int) int {

	if n <= 0 {

		return 1

	}

	return n * Factorial(n-1)

}

// n!/(n-k)!*k!

// Output:
// n == 0: 1
// n == 1: 1 1
// n == 2: 1 2 1
// n == 3: 1 3 3 1
// n == 4: 1 4 6 4 1
// n == 5: 1 5 10 10 5 1
