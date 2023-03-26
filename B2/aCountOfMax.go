package main

import "fmt"

/* STATEMENT
Последовательность состоит из натуральных чисел и завершается числом 0. Всего вводится не более 10000 чисел (не считая завершающего числа 0). Определите, сколько элементов этой последовательности равны ее наибольшему элементу.
Числа, следующие за числом 0, считывать не нужно.
INPUT
Вводится последовательность целых чисел, оканчивающаяся числом 0 (само число 0 в последовательность не входит).
OUTPUT
Выведите ответ на задачу.
EXAMPLE 1
input
1
7
9
0
output
1
EXAMPLE 2
input
1
3
3
1
0
output
2
*/

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
