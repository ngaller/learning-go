package main

import "fmt"

type Sortable interface {
	// Compare returns 0, 1 or -1 if the current element is resp. equal to, bigger or smaller than other
	Compare(other Sortable) int
}

type Heap []Sortable

func (h Heap) Add(elem Sortable) Heap {
	h = append(h, elem)
	h.shiftup(len(h) - 1)
	return h
}

func (h Heap) Pop() (Heap, Sortable) {
	if len(h) == 0 {
		return h, nil
	}
	x := h[0]
	h[0] = h[len(h)-1]
	h = h[:len(h)-1]
	h.shiftdown(0)
	return h, x
}

func (h Heap) shiftup(i int) {
	if i == 0 {
		return
	}
	p := (i - 1) / 2
	// parent is already smaller - OK
	if h[p].Compare(h[i]) <= 0 {
		return
	}
	h[p], h[i] = h[i], h[p]
	h.shiftup(p)
}

func (h Heap) shiftdown(i int) {
	if i*2+1 >= len(h) {
		// already at bottom
		return
	}
	c1, c2 := i*2+1, i*2+2
	if c1 == len(h)-1 || h[c1].Compare(h[c2]) < 0 {
		// check left child (it's either only child, or the smallest)
		if h[i].Compare(h[c1]) > 0 {
			// swap to ensure parent is still the smallest
			h[c1], h[i] = h[i], h[c1]
			h.shiftdown(c1)
		}
	} else {
		// check right
		if h[i].Compare(h[c2]) > 0 {
			// swap to ensure parent is still the smallest
			h[c2], h[i] = h[i], h[c2]
			h.shiftdown(c2)
		}
	}
}

type Order struct {
	t     int
	d     int
	index int
}

func (a Order) Compare(other Sortable) int {
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
	var norder, di, ti int
	var h Heap

	fmt.Scan(&norder)
	for i := 0; i < norder; i++ {
		fmt.Scan(&di, &ti)
		h = h.Add(Order{ti, di, i + 1})
	}
	prev, prevIndex := 0, 0
	for h, s := h.Pop(); s != nil; h, s = h.Pop() {
		o := s.(Order)
		if o.d+o.t < prev {
			msg := fmt.Sprintf("Invalid order: %d vs %d (index %d, %d)",
				o.d+o.t, prev, o.index, prevIndex)
			panic(msg)
		}
		prev = o.d + o.t
		prevIndex = o.index
		fmt.Printf("%d ", o.index)
	}

	//    for _, order := range(h) {
	//        fmt.Printf("%d ", order.(Order).index)
	//    }
}
