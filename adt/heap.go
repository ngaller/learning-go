package adt

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

// func (h Heap) String() string {
// 	return fmt.Sprintf("%d item(s)", len(h))
// }
