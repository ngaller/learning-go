package main

import (
	"fmt"
	"github.com/nicocrm/learning-go/pack"
)

func main() {
	test1 := pack.test1()
	test2 := pack.Test2()
	fmt.Printf("Test1 = %d, Test2 = %s\n", test1, test2)
	// scores := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// test := make([]int, 6)
	// copy(test, scores[:3])
	// // power := 9000
	// // scores[power] = 4
	// fmt.Printf("3 is %d\n", scores[3])
	// fmt.Printf("Length is %d\n", len(scores))
	// fmt.Printf("Cap is %d\n", cap(scores))
	// fmt.Printf("Scores is %v\n", scores)
	// fmt.Printf("Test is %v\n", test)
	// // fmt.Printf("Test2 is %v\n", test2)
}
