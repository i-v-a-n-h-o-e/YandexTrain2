package main

import "fmt"

func main() {
	var length, n, tmp int
	var a []int
	fmt.Scan(&length, &n)
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		a = append(a, tmp)
	}
	flag := length%2 != 0
	for i := 0; i < n; i++ {
		if (a[i] < length/2) && (a[i+1] > length/2) && flag {
			fmt.Printf("%d %d", a[i], a[i+1])
			break
		}
		if (a[i] == length/2) && flag {
			fmt.Print(a[i])
			break
		}
		if (a[i] < length/2) && (a[i+1] >= length/2) && !flag {
			fmt.Printf("%d %d", a[i], a[i+1])
			break
		}
	}
}
