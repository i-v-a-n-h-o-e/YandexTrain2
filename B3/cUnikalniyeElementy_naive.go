package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
STATEMENT

Дан список. Выведите те его элементы, которые встречаются в списке только один раз.
Элементы нужно выводить в том порядке, в котором они встречаются в списке.
INPUT
Вводится список чисел. Все числа списка находятся на одной строке.
OUTPUT
Выведите ответ на задачу.
EXAMPLE 1
input
1 2 2 3 3 3
output
1
EXAMPLE 2
input
4 3 5 2 5 1 3 5
output
4 2 1
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	a := make([]int, 0)
	var c int

	for _, now := range strings.Split(input, " ") {
		numberNow, _ := strconv.Atoi(now)
		a = append(a, numberNow)
	}
	for _, now := range a {
		c = 0
		for _, reimplementation := range a {
			if now == reimplementation {
				c++
				if c > 1 {
					break
				}
			}
		}
		if c == 1 {
			fmt.Printf("%v ", now)
		}
	}
}
