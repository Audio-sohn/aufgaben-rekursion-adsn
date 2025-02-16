package search

import "aufgaben-rekursion/lists"

// FindSorted sucht in einer sortierten Liste nach dem ersten Vorkommen von x.
// Falls x nicht gefunden wird, wird -1 zurückgegeben.
// Da die Liste sortiert ist, wird die binäre Suche verwendet.
func FindSorted(list []int, x int) int {

	// Rekursionsanker: wenn liste leer, dann -1 returnen
	if lists.Empty(list) {

		return -1

	}

	// Rekursionsanker: wenn mittleres element == suchwert
	if list[len(list)/2] == x {

		return len(list) / 2

	}

	// Wenn x größer als die hälfte der liste, untere hälfte discarden
	// und untere stellen auf den index aufsummieren

	if list[len(list)/2] < x {

		return len(list)/2 + FindSorted(list[len(list)/2:], x)

	}

	// Wenn x kleiner als die hälfte der liste, obere hälfte discarden
	// nichts zum index hinzuzählen

	return FindSorted(list[:len(list)/2], x)

}
