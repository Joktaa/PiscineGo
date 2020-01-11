package piscine

func Capitalize(s string) string {
	capital := true
	resultat := ""

	for id := range s {
		if capital == true {
			resultat += ToUpper(string(s[id]))
		} else if capital == false {
			resultat += ToLower(string(s[id]))
		} else {
			resultat += string(s[id])
		}

		if rune(s[id]) > 47 && rune(s[id]) < 58 {
			capital = false
		} else if rune(s[id]) > 64 && rune(s[id]) < 91 {
			capital = false
		} else if rune(s[id]) > 96 && rune(s[id]) < 123 {
			capital = false
		} else {
			capital = true
		}
	}

	return resultat
}
