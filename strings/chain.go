package strings

// Chain erwartet einen String und hängt ihn n mal hintereinander.
// Liefert das Ergebnis.
func Chain(s string, n int) string {

	// rekursionsanker, wenn n drained, dann nichts mehr addieren und
	// rekursion abbrechen
	if n <= 0 {

		return ""

	}

	// ansonsten weiter machen und n dekrementieren
	return s + Chain(s, n-1)
}
