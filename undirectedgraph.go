package graph

import (
	"bytes"
	"container/list"
	"github.com/emirpasic/gods/trees/binaryheap"
	"math"
	"strconv"
)

type UVertex interface {
	Id() string
	Equal(uv UVertex) bool
	Repr() string
	Edges() *list.List
	Clone() UVertex // Clone self except Edges
}

type Edge interface {
	To() UVertex
	From() UVertex
	Weight() float64
}

type UWGraph map[string]UVertex

func NewUWGraph() UWGraph {
	return make(UWGraph)
}

func (g UWGraph) Size() int {
	return len(g)
}

func (g UWGraph) Has(v UVertex) bool {
	var _, ok = g[v.Id()]
	return ok
}

func (g UWGraph) HasBoth(a, b UVertex) bool {
	return g.Has(a) && g.Has(b)
}

func (g UWGraph) Add(v UVertex) {
	if !g.Has(v) {
		g[v.Id()] = v
	}
}

func (g UWGraph) Connect(from, to UVertex, weight float64) error {
	if !g.HasBoth(from, to) {
		return ErrMissingVertex
	}
	from = g[from.Id()]
	to = g[to.Id()]
	from.Edges().PushBack(newEdge(from, to, weight))
	to.Edges().PushBack(newEdge(to, from, weight))
	return nil
}

func (g UWGraph) Connected(from, to UVertex) bool {
	if !g.HasBoth(from, to) {
		return false
	}
	for e := g[from.Id()].Edges().Front(); e != nil; e = e.Next() {
		var v = e.Value.(Edge).To()
		if v.Equal(to) {
			return true
		}
	}
	for e := g[to.Id()].Edges().Front(); e != nil; e = e.Next() {
		var v = e.Value.(Edge).To()
		if v.Equal(from) {
			return true
		}
	}
	return false
}

func (g UWGraph) Disconnect(from, to UVertex) {
	if !g.HasBoth(from, to) {
		return
	}
	var vertices = g[from.Id()].Edges()
	for e := vertices.Front(); e != nil; e = e.Next() {
		var v = e.Value.(Edge).To()
		if v.Equal(to) {
			vertices.Remove(e)
			break
		}
	}
	vertices = g[to.Id()].Edges()
	for e := vertices.Front(); e != nil; e = e.Next() {
		var v = e.Value.(Edge).To()
		if v.Equal(from) {
			vertices.Remove(e)
			break
		}
	}
}

func (g UWGraph) Cyclic() bool {
	var visited = newSet()
	for _, node := range g {
		if !visited.contains(node.Id()) && g.cyclic(node, nil, visited) {
			return true
		}
	}
	return false
}

func (g UWGraph) cyclic(node, parent UVertex, visited set) bool {
	visited.add(node.Id())
	for e := node.Edges().Front(); e != nil; e = e.Next() {
		var next = e.Value.(Edge).To()
		if parent != nil && next.Equal(parent) {
			continue
		}
		if visited.contains(next.Id()) ||
			g.cyclic(next, node, visited) {
			return true
		}
	}
	return false
}

func (g UWGraph) repr() string {
	var buff = &bytes.Buffer{}
	for _, vertex := range g {
		buff.WriteString(vertex.Id())
		buff.WriteString(" => [ ")
		for e := vertex.Edges().Front(); e != nil; e = e.Next() {
			var e = e.Value.(Edge)
			buff.WriteString("<")
			buff.WriteString(e.To().Id())
			buff.WriteString(":")
			buff.WriteString(strconv.FormatFloat(e.Weight(), 'f', 2, 64))
			buff.WriteString("> ")
		}
		buff.WriteString("]\n")
	}
	return buff.String()
}

type row struct {
	previous UVertex
	weight   float64
}

func (g UWGraph) Path(from, to UVertex) (*Path, error) {
	if !g.HasBoth(from, to) {
		return nil, ErrMissingVertex
	}
	var unvisited = newSet()
	var table = make(map[string]*row)
	for id, _ := range g {
		unvisited.add(id)
		table[id] = &row{
			previous: nil,
			weight:   math.Inf(1),
		}
	}
	table[from.Id()].weight = 0
	table[from.Id()].previous = from

	var queue = list.New() // queue of Vertices
	queue.PushBack(g[from.Id()])
	for queue.Len() > 0 {
		var el = queue.Front()
		var node = el.Value.(UVertex)
		if !unvisited.contains(node.Id()) {
			queue.Remove(el)
			continue
		}
		for e := node.Edges().Front(); e != nil; e = e.Next() {
			var edge = e.Value.(Edge)
			var rec = table[edge.To().Id()]
			var dis = table[node.Id()].weight + edge.Weight()
			if rec.weight > dis {
				rec.weight = dis
				rec.previous = node
			}
			if unvisited.contains(edge.To().Id()) {
				queue.PushBack(edge.To())
			}
		}
		unvisited.remove(node.Id())
		queue.Remove(el)
	}
	return newPath(table, g[to.Id()]), nil
}

func (g UWGraph) randomVertex() (UVertex, bool) {
	for _, v := range g {
		return v, true
	}
	return nil, false
}

// MinTree return minimum spanning tree
// using Prim's algorithm
func (g UWGraph) MinTree() UWGraph {
	var res = NewUWGraph()
	var v, ok = g.randomVertex()
	if !ok {
		return res
	}
	var queue = binaryheap.NewWith(func(a, b interface{}) int {
		var weightA = a.(Edge).Weight()
		var weightB = b.(Edge).Weight()
		if weightA < weightB {
			return -1
		}
		if weightA > weightB {
			return 1
		}
		return 0
	})
	for e := v.Edges().Front(); e != nil; e = e.Next() {
		queue.Push(e.Value)
	}
	res.Add(v.Clone())
	for res.Size() < g.Size() {
		var e, ok = queue.Pop()
		if !ok {
			return res
		}
		var edge = e.(Edge)
		if res.Has(edge.To()) {
			continue
		}
		res.Add(edge.To().Clone())
		res.Connect(edge.From(), edge.To(), edge.Weight())
		for e := edge.To().Edges().Front(); e != nil; e = e.Next() {
			queue.Push(e.Value)
		}
	}
	return res
}
