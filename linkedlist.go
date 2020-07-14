package graph

type node struct {
	val  Vertex
	next *node
}

func (n *node) equal(other *node) bool {
	return n.val.Equal(other.val)
}

func (n *node) string() string {
	return n.val.Repr()
}

type LinkedList struct {
	head *node
	tail *node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Append(v Vertex) {
	var node = &node{val: v}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}
}

func (l *LinkedList) Remove(v Vertex) bool {
	if l.head == nil {
		return false
	}
	if l.head.val.Equal(v) {
		l.head = l.head.next
		if l.head == nil {
			l.tail = nil
		}
		return true
	}
	var n = l.head
	for n.next != nil && !n.next.val.Equal(v) {
		n = n.next
	}
	if n.next == l.tail {
		l.tail = n
	}
	if n.next != nil {
		n.next = n.next.next
		return true
	}
	return false
}

func (l *LinkedList) Iterator() <-chan Vertex {
	var ch = make(chan Vertex, 10)
	go func() {
		defer close(ch)
		var node = l.head
		for node != nil {
			ch <- node.val
			node = node.next
		}
	}()
	return ch
}