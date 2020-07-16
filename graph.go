package graph

import (
	"bytes"
	"container/list"
	"errors"
)

var ErrMissingVertex = errors.New("vertex is missing")
var ErrCyclicGraph = errors.New("cyclic graph")

type Vertex interface {
	Id() string
	Repr() string
	Equal(v Vertex) bool
}

type DiGraph map[Vertex]*LinkedList

func NewDiGraph() DiGraph {
	return make(DiGraph)
}

func (g DiGraph) Has(v Vertex) bool {
	var _, ok = g[v]
	return ok
}

func (g DiGraph) Edges(v Vertex) (*LinkedList, error) {
	if !g.Has(v) {
		return nil, ErrMissingVertex
	}
	return g[v], nil
}

func (g DiGraph) Add(v Vertex) *LinkedList {
	if g.Has(v) {
		return g[v]
	}
	var ll = NewLinkedList()
	g[v] = ll
	return ll
}

func (g DiGraph) Remove(v Vertex) {
	if !g.Has(v) {
		return
	}
	delete(g, v)
	for _, ll := range g {
		ll.Remove(v)
	}
}

func (g DiGraph) Connect(from, to Vertex) error {
	if !g.Has(from) || !g.Has(to) {
		return ErrMissingVertex
	}
	g[from].Append(to)
	return nil
}

func (g DiGraph) Disconnect(from, to Vertex) error {
	if !g.Has(from) || !g.Has(to) {
		return ErrMissingVertex
	}
	g[from].Remove(to)
	return nil
}

func (g DiGraph) repr() string {
	var buff = &bytes.Buffer{}
	for vertex, ll := range g {
		buff.WriteString(vertex.Id())
		buff.WriteString(" is connected with [ ")
		for v := range ll.Iterator() {
			buff.WriteString(v.Id())
			buff.WriteString(" ")
		}
		buff.WriteString("]\n")
	}
	return buff.String()
}

func (g DiGraph) DFS(v Vertex) <-chan Vertex {
	var out = make(chan Vertex, 10)
	go g.traverse(v, out, false)
	return out
}

func (g DiGraph) BFS(v Vertex) <-chan Vertex {
	var out = make(chan Vertex, 10)
	go g.traverse(v, out, true)
	return out
}

func (g DiGraph) traverse(v Vertex, out chan<- Vertex, useQueue bool) {
	defer close(out)
	if !g.Has(v) {
		return
	}
	var visited = newSet()
	var ll = list.New()
	ll.PushBack(v)
	for ll.Len() > 0 {
		var el *list.Element
		if useQueue {
			el = ll.Front()
		} else {
			el = ll.Back()
		}
		var cv = el.Value.(Vertex)
		ll.Remove(el)
		out <- cv
		visited.add(cv)
		for nv := range g[cv].Iterator() {
			if !visited.contains(nv) {
				ll.PushBack(nv)
			}
		}
	}
}

// Sorted implements topological sorting on directed acyclic graph (DAG)
func (g DiGraph) Sorted() ([]Vertex, error) {
	if g.Cyclic() {
		return make([]Vertex, 0), ErrCyclicGraph
	}
	var out = list.New()
	var visited = newSet()
	for parent, _ := range g {
		g.sorted(parent, visited, out)
	}
	var result = make([]Vertex, out.Len())
	for out.Len() > 0 {
		result[out.Len()-1] = out.Back().Value.(Vertex)
		out.Remove(out.Back())
	}
	return result, nil
}

func (g DiGraph) sorted(v Vertex, visited set, out *list.List) {
	if !g.Has(v) || visited.contains(v) {
		return
	}
	var node = g[v].head
	for node != nil {
		g.sorted(node.val, visited, out)
		node = node.next
	}
	out.PushFront(v)
	visited.add(v)
}

func (g DiGraph) Cyclic() bool {
	var visited = newSet()
	var visiting = newSet()
	for parent, _ := range g {
		if g.cyclic(parent, visiting, visited) {
			return true
		}
	}
	return false
}

func (g DiGraph) cyclic(v Vertex, visiting, visited set) bool {
	if visiting.contains(v) {
		return true
	}
	if !g.Has(v) || visited.contains(v) {
		return false
	}
	visiting.add(v)
	var node = g[v].head
	for node != nil {
		if g.cyclic(node.val, visiting, visited) {
			return true
		}
		node = node.next
	}
	visited.add(v)
	visiting.remove(v)
	return false
}
