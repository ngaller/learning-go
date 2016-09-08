package main

import "fmt"

func main() {
	var s string

	fmt.Scan(&s)
	fmt.Printf("Read: %s", s)
	var n int
	_, err := fmt.Scan(&n)
	fmt.Printf("Read: %d", n)
	fmt.Printf("Err: %v", err)

}
