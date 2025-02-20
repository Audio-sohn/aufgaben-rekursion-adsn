package strings

// Contains prüft, ob der String s die Sequenz seq enthält.
func Contains(s, seq string) bool {

	// FALSCH: Immer volle sequenz vergleichen!

	// Rekursionsanker: Wenn s < seq , false returnen
	if Length(s) < Length(seq) {

		return false

	}

	// Rekursionsanker: Wenn der Anfang von s mit seq übereinstimmt, true returnen

	if s[:Length(seq)] == seq {

		return true

	}

	return Contains(s[1:], seq)

}
