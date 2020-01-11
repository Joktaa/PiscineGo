package piscine

func ToLower(s string) string {
	resultat := ""

	for id := range s {
		if rune(s[id]) > 64 && rune(s[id]) < 91 {
			resultat += string(s[id] + 32)
		} else {
			resultat += string(s[id])
		}
	}

	return resultat
}
