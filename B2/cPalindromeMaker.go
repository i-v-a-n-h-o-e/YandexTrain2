package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	n := 0
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			n++
		}
	}
	fmt.Print(n)
}
