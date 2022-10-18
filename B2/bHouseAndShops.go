package main

import (
	"fmt"
)

func main() {
	var houses []int
	var shops []int
	var tmp, start, end, min, dist1, dist2 int
	for i := 0; i < 10; i++ {
		fmt.Scan(&tmp)

		if tmp == 1 {
			houses = append(houses, i)
		} else if tmp == 2 {
			shops = append(shops, i)
		}
	}
	max := -1

	for i := 0; i < len(houses); i++ {
		start = 0
		end = len(shops) - 1
		for (start != end) && ((start + 1) != end) && (shops[start] < houses[i]) && (shops[end] > houses[i]) {
			tmp = shops[(start+end)/2]
			if tmp < houses[i] {
				start = (start + end) / 2
			} else {
				end = (start + end) / 2
			}
		}

		if houses[i]-shops[start] > 0 {
			dist1 = houses[i] - shops[start]
		} else {
			dist1 = shops[start] - houses[i]
		}

		if shops[end]-houses[i] > 0 {
			dist2 = shops[end] - houses[i]
		} else {
			dist2 = houses[i] - shops[end]
		}

		if dist1 < dist2 {
			min = dist1
		} else {
			min = dist2
		}

		if min > max {
			max = min
		}
	}
	fmt.Print(max)
}
