package list

type List *Node
type Node struct {
	content string
	next    *Node
}

func Length(l List) int {
	curs := l
	c := 0
	for curs != nil {
		c++
		curs = curs.next
	}
}
func AddFront(l List, x string) list {
	n := new(Node)
	n.content = x
	n.next = l
	return n
}
func AddInOrder(l List, x string) list {
	n := new(Node)
	n.content = x
	curs := l
	prev := nil
	for curs != nil && curs.content < x {
		prev = curs
		curs = curs.next
	}
	n.next = curs
	prev.next = n
	if prev == nil {
		return n
	} else {
		prev.next = n
		return l
	}
}
func Concatenate(first, second list) list {
	if first == nil { //Senza questo if rischiamo un panic in curs.next
		return second
	}
	curs := first
	for curs.next != nil {
		curs = curs.next
	}
	curs.next = second
	return first
}
func String(l List) string {

}
