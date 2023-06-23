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
type Set struct {
	array            [][]int
	numberOfElements int
	sizeOf           int
}

func hash(a int, sizeOf int) int {
	return a % sizeOf
}

func (s *Set) find(element int) bool {
	for _, now := range s.array[hash(element, s.sizeOf)] {
		if element == now {
			return true
		}
	}
	return false
}
func (s *Set) add(newElement int) {
	if s.sizeOf <= s.numberOfElements {
		s.sizeOf = s.sizeOf * 2
		tempArray := make([][]int, s.sizeOf)
		for index := 0; index < s.sizeOf/2; index++ {
			for _, nowElement := range s.array[index] {
				tempArray[hash(nowElement, s.sizeOf)] = append(tempArray[hash(nowElement, s.sizeOf)], nowElement)
			}
		}
		s.array = tempArray
	}
	if !s.find(newElement) {
		s.array[hash(newElement, s.sizeOf)] = append(s.array[hash(newElement, s.sizeOf)], newElement)
		s.numberOfElements++
	}
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	all := Set{
		make([][]int, 10),
		0,
		10,
	}
	Repeater := Set{
		make([][]int, 10),
		0,
		10,
	}

	a := make([]int, 0)

	for _, now := range strings.Split(input, " ") {
		numberNow, _ := strconv.Atoi(now)
		a = append(a, numberNow)
		if all.find(numberNow) {
			Repeater.add(numberNow)
		} else {
			all.add(numberNow)
		}
	}

	for _, now := range a {
		if !Repeater.find(now) {
			fmt.Printf("%d ", now)
		}
	}
}
