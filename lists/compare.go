package lists

// Liefert true, falls die beiden Listen gleich sind.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func ListsEqual(list1, list2 []int) bool {
	// TODO

	// wenn nur eine der listen leer, false returnen
	if (Empty(list1) || Empty(list2)) && !(Empty(list1) && Empty(list2)) {

		return false

		// wenn beide listen leer, true returnen
	} else if Empty(list1) && Empty(list2) {

		return true

		// wenn elemente in erster stelle ungleich, false returnen
	} else if list1[0] != list2[0] {

		return false

	}

	return ListsEqual(list1[1:], list2[1:])
}
