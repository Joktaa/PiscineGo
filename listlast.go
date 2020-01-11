package piscine

func ListLast(l *List) interface{} {
	var result interface{}
	head := l.Head

	for head != nil {
		result = head.Data
		head = head.Next
	}

	return result
}
