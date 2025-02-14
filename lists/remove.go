package lists

// Liefert eine Liste, die alle Elemente aus list enthält,
// außer dem an Stelle pos.
// Wenn pos außerhalb des Bereichs der Liste liegt, wird die
// ursprüngliche Liste zurückgegeben.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func RemoveElement(list []int, pos int) []int {

	// prüfe ob liste empty und position abgelaufen
	// falls ja, erstes element (element an pos) wegschneiden und restliche liste returnen
	if pos == 0 {

		return list[1:]

	}

	// wenn liste empty und position nicht abgelaufen, leere liste returnen
	// dann ist ja die liste durchgesucht und pos außerhalb von der liste
	if Empty(list) && pos != 0 {

		return nil

	}

	// wenn liste nicht empty und position nicht abgelaufen, in nächste runde gehen
	// bei nächstem call dann erstes element "wegfressen" und position dekrementieren
	return append(list[:1], RemoveElement(list[1:], pos-1)...)

}
