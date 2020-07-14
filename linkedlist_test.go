package graph

import (
	"strconv"
	"testing"
)

func Mock(names []string) *LinkedList {
	var ll = NewLinkedList()
	ll.Append(&vertex{id: names[0]})
	for i := 1; i < len(names); i++ {
		ll.Append(&vertex{id: names[i]})
	}
	return ll
}

func TestLinkedList_Append(t *testing.T) {
	var names = []string{"A", "B", "C", "D", "E", "F"}
	var ll = Mock(names)
	var node = ll.head
	var i int
	for node != nil {
		if node.val.Id() != names[i] {
			t.Errorf("Order Violation: %s != %s", node.val.Id(), names[i])
		}
		node = node.next
		i++
	}
	if ll.tail.val.Id() != names[len(names)-1] {
		t.Errorf("Wrong tail")
		t.FailNow()
	}
	for ll.tail != nil {
		ll.Remove(ll.tail.val)
	}
	ll.Append(&vertex{id: "head"})
	if ll.head == nil {
		t.Error("Head is not set")
		t.FailNow()
	}
	if ll.tail == nil {
		t.Error("Tail is not set")
		t.FailNow()
	}
	if ll.head != ll.tail {
		t.Errorf("Head != Tail: %s, %s", ll.head.string(), ll.tail.string())
		t.FailNow()
	}

	var ll2 = Mock([]string{"TEST"})
	ll2.Remove(&vertex{id: "TEST"})
	if ll2.tail != nil && ll2.head != nil {
		t.Fatal("Remove failed")
	}
	ll2.Append(&vertex{id: "A"})
	if ll2.head != ll2.tail || ll2.tail == nil {
		t.Error("New node has not been set")
	}
}

func TestLinkedList_Remove(t *testing.T) {
	var names = []string{"A", "B", "C", "D", "E", "F"}
	var ll = Mock(names)
	var tail = *ll.tail
	ll.Remove(tail.val)
	if tail.equal(ll.tail) {
		t.Errorf("Tail has not been removed: %s == %s", ll.tail.string(), tail.string())
		t.FailNow()
	}
	var head = *ll.head
	ll.Remove(head.val)
	if head.equal(ll.head) {
		t.Errorf("Head has not been removed: %s == %s", head.string(), ll.head.string())
		t.FailNow()
	}
	var mid = &vertex{id: "D"}
	ll.Remove(mid)
	var node = ll.head
	for node != nil {
		if node.val.Equal(mid) {
			t.Errorf("Node has not been removed: %s", node.string())
			t.FailNow()
		}
		node = node.next
	}
	for ll.tail != nil {
		ll.Remove(ll.tail.val)
	}
	if ll.head != nil {
		t.Errorf("Head has not been removed: %s", ll.head.string())
	}
}

func TestLinkedList_Iterator(t *testing.T) {
	var ll = NewLinkedList()
	for i := 0; i < 100; i++ {
		ll.Append(&vertex{id: strconv.Itoa(i)})
	}
	var j int
	var c = ll.Iterator()
	for v := range c {
		if n, _ := strconv.Atoi(v.Id()); n != j {
			t.Errorf("Wrong vertex received, expected: %v, got: %v", j, n)
		}
		j++
	}
}
