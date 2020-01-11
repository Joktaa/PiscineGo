package piscine

func SortIntegerTable(table []int) {
	var max int
	var id int
	tmp := make([]int, len(table))
	tmp = table
	resultat := make([]int, len(table))

	for i := 0; i < len(table); i++ {
		max = 0
		id = 0
		for j := 0; j < len(tmp); j++ {
			if tmp[j] > max {
				max = tmp[j]
				id = j
			}
		}
		resultat[i] = max
		tmp = remove(tmp, id)
	}
	table = resultat
}

func remove(s []int, i int) []int {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
