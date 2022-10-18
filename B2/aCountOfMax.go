package main

import "fmt"

func main() {
	var current int
	fmt.Scanf("%d", &current)
	maxNumber := current
	maxCount := 0

	for current != 0 {
		if current > maxNumber {
			maxNumber = current
			maxCount = 1
		} else if current == maxNumber {
			maxCount++
		}
		fmt.Scanf("%d", &current)
	}
	fmt.Print(maxCount)
}
