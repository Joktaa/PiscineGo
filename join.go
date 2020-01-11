package piscine

func Join(strs []string, sep string) string {
	resultat := ""

	for id, val := range strs {
		if id != 0 {
			resultat += sep
		}
		resultat += val
	}
	return resultat
}
