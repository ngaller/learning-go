package main

import "fmt"

func main() {
	var n int
	//var format string
	n = 6
	staircase := "#"
	//fmt.Scan(&n)
	for i := 0; i < n; i++ {
		format := fmt.Sprintf("%%%ds\n", n)
		fmt.Printf(format, staircase)
		staircase += "#"
	}
}
