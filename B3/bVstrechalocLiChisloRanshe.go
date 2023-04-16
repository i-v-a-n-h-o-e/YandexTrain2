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

Во входной строке записана последовательность чисел через пробел.
Для каждого числа выведите слово YES (в отдельной строке),
если это число ранее встречалось в последовательности или NO, если не встречалось.
INPUT
Вводится список чисел. Все числа списка находятся на одной строке.
OUTPUT
Выведите ответ на задачу.
EXAMPLE
input
1 2 3 2 3 4
output
NO
NO
NO
YES
YES
NO
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
	s.array[hash(newElement, s.sizeOf)] = append(s.array[hash(newElement, s.sizeOf)], newElement)
	s.numberOfElements++
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	set := Set{
		make([][]int, 10),
		0,
		10,
	}

	for _, now := range strings.Split(input, " ") {
		numberNow, _ := strconv.Atoi(now)
		if set.find(numberNow) {
			fmt.Println("YES")
		} else {
			set.add(numberNow)
			fmt.Println("NO")
		}
	}
}
