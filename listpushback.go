package piscine

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = n
	} else {
		currentNode := l.Head
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = n
	}
}
