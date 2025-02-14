package lists

// Liefert zwei Listen:
// - Eine, die alle Elemente aus list enthält, die kleiner oder gleich key sind.
// - Eine, die alle übrigen Elemente aus list enthält.
func Partition(list []int, key int) ([]int, []int) {

	// wenn die liste leer ist, "leer" returnen und rekursion abreißen

	// Verwende Kopien von list, damit die ursprüngliche Liste nicht verändert wird.
	// l1 := append([]int{}, list...)
	// l2 := append([]int{}, list...)

	if Empty(list) {

		return l1, l2

	}

	// wenn aktuelles element größer als key, löschen und in andere liste schreiben
	if list[0] > key {

		l1 = append(l1, list[0])

	} else {

		l2 = append(l2, list[0])

	}

	return Partition(list[1:], key)
}
