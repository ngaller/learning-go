package main

import (
	"fmt"
)

func dist(s string) int {
	n := 0
	r := []rune(s)
	l := len(r)
	for i := 0; i < l/2; i++ {
		op := int(r[i] - r[l-1-i])
		if op < 0 {
			op = -op
		}
		n += op
	}
	return n
}

func main() {
	var n int
	var s string

	fmt.Scan(&n)
	for ; n > 0; n-- {
		fmt.Scan(&s)
		fmt.Printf("%d\n", dist(s))
	}
}
