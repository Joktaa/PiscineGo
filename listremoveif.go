package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	head := l.Head
	res := &List{}
	empty := true

	for head != nil {
		if head.Data != data_ref {
			ListPushBack(res, head.Data)
			empty = false
		}
		head = head.Next
	}
	if empty {
		l.Head = nil
	} else {
		l.Head.Data = res.Head.Data
		l.Head.Next = res.Head.Next
	}
}
