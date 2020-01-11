package piscine

func BasicJoin(strs []string) string {
	resultat := ""

	for _, val := range strs {
		resultat += val
	}

	return resultat
}
