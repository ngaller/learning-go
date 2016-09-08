package main

import "fmt"

func main() {
	source, k := "39282", 2
	//source, k := "3943", 1

	s := []byte(source)
	n := len(s)
	changed := make([]bool, n)

	for i := 0; i < n/2; i++ {
		if s[i] < s[n-i-1] {
			s[i] = s[n-i-1]
			changed[i] = true
			fmt.Printf("Changed character at %d\n", i)
			k--
		}
		if s[i] > s[n-i-1] {
			s[n-i-1] = s[i]
			changed[n-1-i] = true
			fmt.Printf("Changed character at %d\n", n-1-i)
			k--
		}
	}

	for i := 0; i < (n+1)/2 && k > 0; i++ {
		if s[i] < '9' {
			if changed[i] || changed[n-i-1] || i == n-1-i {
				s[i] = '9'
				s[n-1-i] = '9'
				k--
			} else if k >= 2 {
				s[i] = '9'
				s[n-1-i] = '9'
				k -= 2
			}
		}
	}

	if k < 0 {
		fmt.Printf("-1\n")
	} else {
		fmt.Printf("%s\n", s)
	}

}
