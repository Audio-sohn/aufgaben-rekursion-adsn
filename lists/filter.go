package lists

// Liefert eine Liste mit allen Elementen aus list, die kleiner oder gleich key sind.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func FilterLess(list []int, key int) []int {
	// Gehen Sie ähnlich wie bei Remove vor:
	// Wenn die Liste leer ist, ist das Ergebnis die leere Liste.

	if Empty(list) {

		return list

	}

	// Wenn das erste Element größer als key ist, ist das Ergebnis die gefilterte Restliste.

	if list[0] > key {

		return FilterLess(list[1:], key)

	}

	// Wenn das erste Element kleiner oder gleich key ist, ist das Ergebnis das erste Element
	// plus die gefilterte Restliste.
	if list[0] < key {

		return append([]int{list[0]}, FilterLess(list[1:], key)...)

	}

	return list
}

// Liefert eine Liste mit allen Elementen aus list, die echt größer als key sind.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func FilterGreater(list []int, key int) []int {

	if Empty(list) {

		return list

	}

	if list[0] < key {

		return FilterGreater(list[1:], key)

	}

	return append([]int{list[0]}, FilterGreater(list[1:], key)...)
}
