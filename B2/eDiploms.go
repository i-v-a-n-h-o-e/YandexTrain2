package main

import "fmt"

func main() {
	max := 0
	sum := 0
	var n, c int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&c)
		sum = sum + c
		if max < c {
			max = c
		}
	}
	fmt.Print(sum - max)
}
