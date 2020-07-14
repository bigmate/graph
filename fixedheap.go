package graph

import (
	"bytes"
	"errors"
	"strconv"
)

var ErrHeapOverflow = errors.New("heap overflow")
var ErrNoValue = errors.New("no value to return")

type FixedHeap struct {
	arr []Edge
	len int
	cap int
}

func NewFixedHeap(size int) *FixedHeap {
	return &FixedHeap{
		arr: make([]Edge, size+1),
		cap: size,
	}
}

func (p *FixedHeap) less(i, j int) bool {
	return p.arr[i].Weight() < p.arr[j].Weight()
}

func (p *FixedHeap) swap(i, j int) {
	p.arr[i], p.arr[j] = p.arr[j], p.arr[i]
}

func (p *FixedHeap) swim(i int) {
	for i > 1 && p.less(i, i/2) {
		p.swap(i, i/2)
		i /= 2
	}
}

func (p *FixedHeap) sink(i int) {
	var left, right = i * 2, i*2 + 1
	for i < p.len {
		var min int
		if left <= p.len && right <= p.len {
			min = p.min(left, right)
		} else if left <= p.len {
			min = left
		} else {
			break
		}
		if p.less(min, i) {
			p.swap(min, i)
			i = min
			left, right = i*2, i*2+1
		} else {
			break
		}
	}
}

func (p *FixedHeap) min(left, right int) int {
	if p.less(left, right) {
		return left
	}
	return right
}

func (p *FixedHeap) repr() string {
	var buff = bytes.Buffer{}
	buff.WriteString("[ ")
	for i := 1; i <= p.Len(); i++ {
		var edge = p.arr[i]
		buff.WriteString(edge.To().Id())
		buff.WriteString(":")
		buff.WriteString(strconv.FormatFloat(edge.Weight(), 'f', 2, 64))
		buff.WriteString(" ")
	}
	buff.WriteString("]")
	return buff.String()
}

func (p *FixedHeap) Len() int {
	return p.len
}

func (p *FixedHeap) Push(e Edge) {
	if p.cap <= p.len {
		panic(ErrHeapOverflow)
	}
	p.len++
	p.arr[p.len] = e
	p.swim(p.len)
}

func (p *FixedHeap) Pop() Edge {
	if p.len <= 0 {
		panic(ErrNoValue)
	}
	if p.len == 1 {
		p.len--
		return p.arr[1]
	}
	var front = p.arr[1]
	p.swap(1, p.len)
	p.len--
	p.sink(1)
	return front
}
