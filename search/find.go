package search

import "aufgaben-rekursion/lists"

// Find sucht in einer Liste nach dem ersten Vorkommen von x
// und gibt dessen Index zurück. Falls x nicht gefunden wird,
// wird -1 zurückgegeben.
func Find(list []int, x int) int {

	// rekurstionsanker: wenn liste drained, -1 returnen
	if lists.Empty(list) {

		return -1

	}

	// rekursionsanker: wenn element gefunden, index nicht inkrementieren
	// und rekursion abbrechen
	if list[0] == x {

		return 0

	}

	return 1 + Find(list[1:], x)
}
