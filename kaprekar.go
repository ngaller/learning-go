package main

import "fmt"
import "strconv"

func main() {
	s := "12345"

	// fmt.Printf("%d\n", len(s)/2)
	// fmt.Printf("%s\n", s[0:2])
	a, err := strconv.Atoi(s[0 : (len(s)+1)/2])
	b, err := strconv.Atoi(s[(len(s)+1)/2:])
	if err != nil {
		panic("Could not convert")
	}
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
}
