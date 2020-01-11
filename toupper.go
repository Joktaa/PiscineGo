package piscine

func ToUpper(s string) string {
	resultat := ""

	for id := range s {
		if rune(s[id]) > 96 && rune(s[id]) < 123 {
			resultat += string(s[id] - 32)
		} else {
			resultat += string(s[id])
		}
	}

	return resultat
}
