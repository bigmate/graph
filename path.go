package graph

import (
	"container/list"
)

type Path struct {
	weight   float64
	vertices *list.List
}

func (p *Path) Weight() float64 {
	return p.weight
}

func (p *Path) Vertices() *list.List {
	return p.vertices
}

func newPath(table map[string]*row, to UVertex) *Path {
	var path = &Path{
		weight:   table[to.Id()].weight,
		vertices: list.New(),
	}
	path.vertices.PushFront(to)
	var current = to
	var previous = table[to.Id()].previous
	for previous != nil && !current.Equal(previous) {
		path.vertices.PushFront(previous)
		current = previous
		previous = table[previous.Id()].previous
	}
	return path
}
