package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("What is yer name?")
	fmt.Scan(&name)
	fmt.Printf("Hello then, %s\n", name)
}
