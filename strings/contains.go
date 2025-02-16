package strings

// Contains prüft, ob der String s die Sequenz seq enthält.
func Contains(s, seq string) bool {

	// Rekursionsanker: Wenn seq drained, true returnen
	// kann nur passieren wenn seq in s auftaucht

	if seq == "" {

		return true

	}

	// Rekursionsanker: Wenn s drained und seq nicht, false returnen
	// kann an dieser stelle nur passieren wenn seq nicht autaucht

	if s == "" {

		return false

	}

	// wenn erstes element von s UNGLEICH erstes element von seq
	// weitermachen ohne erstes element von s

	if s[0] != seq[0] {

		return Contains(s[1:], seq)

	}

	// wenn erstes element von s GLEICH erstes element von seq
	// weiter machen ohne erstes element von s UND ohne erstes element von seq

	return Contains(s[1:], seq[1:])

}
