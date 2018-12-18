package list

//List is what this package is about
type List *Node

//Node is an element of List
type Node struct {
	content string
	next    *Node
}

//Lenght return the lenghts of the list
func Lenght(l List) int {
	curs := l
	c := 0
	for curs != nil {
		c++
		curs = curs.next
	}
	return c
}

//AddFront adds a new item in front of the list
func AddFront(l List, x string) List {
	n := new(Node)
	n.content = x
	n.next = l
	return n
}

//AddInOrder adds a new item to the list keeping alfabetic order
func AddInOrder(l List, x string) List {
	n := new(Node)
	n.content = x
	curs := l
	var prev List
	for curs != nil && curs.content < x {
		prev = curs
		curs = curs.next
	}
	n.next = curs
	prev.next = n
	if prev == nil {
		return n
	}
	prev.next = n
	return l

}

//Concatenate creates a new list made up of the 2 lists one afther the other
func Concatenate(first, second List) List {
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
