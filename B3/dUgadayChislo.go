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

Август и Беатриса играют в игру. Август загадал натуральное число от 1 до n.
Беатриса пытается угадать это число, для этого она называет некоторые множества натуральных чисел.
Август отвечает Беатрисе YES, если среди названных ей чисел есть задуманное или NO в противном случае.
После нескольких заданных вопросов Беатриса запуталась в том, какие вопросы она задавала и
какие ответы получила и просит вас помочь ей определить, какие числа мог задумать Август.
INPUT
Первая строка входных данных содержит число n — наибольшее число, которое мог загадать Август.
Далее идут строки, содержащие вопросы Беатрисы. Каждая строка представляет собой набор чисел, разделенных пробелами.
После каждой строки с вопросом идет ответ Августа: YES или NO. Наконец, последняя строка входных данных содержит одно слово HELP.
OUTPUT
Вы должны вывести (через пробел, в порядке возрастания) все числа, которые мог задумать Август.
EXAMPLE 1
input
10
1 2 3 4 5
YES
2 4 6 8 10
NO
HELP
output
1 3 5
EXAMPLE 2
input
10
1 2 3 4 5 6 7 8 9 10
YES
1
NO
2
NO
3
NO
4
NO
6
NO
7
NO
8
NO
9
NO
10
NO
HELP
output
5
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
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Scan(&n)

	yes := Set{
		make([][]int, n),
		0,
		n,
	}
	for i := 1; i <= n; i++ {
		yes.add(i)
	}
	no := Set{
		make([][]int, n),
		0,
		n,
	}
	a := make([]int, 0)
	var input string
	for input != "HELP" {
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "YES" {
			newYes := Set{
				make([][]int, n),
				0,
				n,
			}
			for _, now := range a {
				if yes.find(now) {
					newYes.add(now)
				}
			}
			yes = newYes
		} else if input == "NO" {
			for _, now := range a {
				no.add(now)
			}
		} else {
			a = make([]int, 0)

			for _, now := range strings.Split(input, " ") {
				numberNow, _ := strconv.Atoi(now)
				a = append(a, numberNow)
			}
		}
	}
	for i := 1; i <= n; i++ {
		if yes.find(i) && !no.find(i) {
			fmt.Printf("%d ", i)
		}
	}
}
