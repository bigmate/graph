package graph

import "container/list"

type set map[interface{}]bool

func newSet() set {
	return make(set)
}

func (s set) contains(val interface{}) bool {
	var _, ok = s[val]
	return ok
}

func (s set) add(val interface{}) {
	s[val] = true
}

func (s set) remove(val interface{}) {
	delete(s, val)
}

type vertex struct {
	id string
}

func (v vertex) Id() string {
	return v.id
}

func (v vertex) Repr() string {
	return v.id
}

func (v vertex) Equal(other Vertex) bool {
	return v.id == other.Id()
}

type uv struct {
	id string
	vs *list.List
}

func (u *uv) Equal(uv UVertex) bool {
	return u.Id() == uv.Id()
}

func (u *uv) Id() string {
	return u.id
}

func (u *uv) Repr() string {
	return "<V: " + u.id + ">"
}

func (u *uv) Edges() *list.List {
	return u.vs
}

func (u *uv) Clone() UVertex {
	return newUV(u.id)
}

func newUV(id string) UVertex {
	return &uv{
		id: id,
		vs: list.New(),
	}
}

type edge struct {
	from UVertex
	to   UVertex
	w    float64
}

func (e *edge) From() UVertex {
	return e.from
}

func (e *edge) To() UVertex {
	return e.to
}

func (e *edge) Weight() float64 {
	return e.w
}

func newEdge(from, to UVertex, w float64) *edge {
	return &edge{
		from: from,
		to:   to,
		w:    w,
	}
}
