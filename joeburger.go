package main

import (
	"fmt"
	"github.com/nicocrm/learning-go/adt"
)

type Order struct {
	t     int
	d     int
	index int
}

func (a Order) Compare(other adt.Sortable) int {
	b, ok := other.(Order)
	if !ok {
		panic("Invalid object type passed to Compare")
	}
	if a.t+a.d > b.t+b.d {
		return 1
	} else if a.t+a.d < b.t+b.d {
		return -1
	}
	return a.index - b.index
}

func main() {
	var h adt.Heap

	h = h.Add(Order{8, 1, 2})
	h = h.Add(Order{8, 1, 1})
	h = h.Add(Order{3, 6, 3})
	h = h.Add(Order{3, 1, 4})
	h = h.Add(Order{4, 3, 5})
	h = h.Add(Order{4, 3, 6})
	h = h.Add(Order{4, 3, 7})
	h = h.Add(Order{4, 3, 8})
	h = h.Add(Order{4, 3, 9})
	h = h.Add(Order{4, 2, 10})
	fmt.Printf("Heap %v", h)
	for h, s := h.Pop(); s != nil; h, s = h.Pop() {
		o := s.(Order)
		fmt.Printf("Order %v\n", o)
	}
}
