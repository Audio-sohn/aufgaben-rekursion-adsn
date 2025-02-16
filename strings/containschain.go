package strings

// ContainsChain liefert true, falls s eine Kette von count aufeinanderfolgenden
// Vorkommen von symbol enthält.
func ContainsChain(s, symbol string, count int) bool {

	// wenn symbol leer, dann true returnen
	if symbol == "" {

		return true

	}

	// Rekursionsanker: Wenn count drained, dann condition erfüllt
	if count <= 0 {

		return true

	}

	// Rekursionsanker: Wenn liste an dieser stelle leer, dann condition nicht erfüllt
	if s == "" {

		return false

	}

	// Rekursionsschritt: Wenn characters gleich, count dekrementieren und
	// string kürzen

	if string(s[0]) == symbol {

		return ContainsChain(s[1:], symbol, count-1)
	}

	// Rekursionsschritt: Wenn characters NICHT gleich, count inkrementieren
	// und string kürzen

	return ContainsChain(s[1:], symbol, count+1)
}

// TODO

// Test ist zwar ok, aber für containschain( "aaaaaaaabb", "b", 2) würde es nicht funktionieren....
