package main

import "fmt"

/*
STATEMENT
В этом году Иван заканчивает школу и поступает в вуз. За время своей учебы он часто участвовал в олимпиадах по информатике и у него накопилось много дипломов. Иван раскладывал дипломы по папкам совершенно бессистемно, то есть любой диплом мог оказаться в любой из папок. К счастью, Иван помнит, сколько дипломов лежит в каждой из папок.
Иван хочет принести в приемную комиссию выбранного вуза папку, в которой находится диплом Московской олимпиады по программированию (такой диплом у Ивана ровно один). Для того чтобы понять, что в данной папке нужного диплома нет, Ивану нужно просмотреть все дипломы из этой папки. Просмотр одного диплома занимает у него ровно одну секунду и он может мгновенно переходить к просмотру следующей папки после окончания просмотра предыдущей. Порядок просмотра папок Иван может выбирать.
По заданному количеству дипломов в каждой из папок требуется определить, за какое наименьшее время в худшем случае Иван поймет, в какой папке содержится нужный ему диплом.
INPUT
В первой строке входного файла записано целое число N (1 ≤ N ≤ 100) - количество папок. Во второй строке записаны N целых чисел a1, a2, ..., aN (1 ≤ ai ≤ 100) - количество дипломов в каждой из папок.
OUTPUT
Выведите одно число - минимальное количество секунд, необходимое Ивану в худшем случае для определения того, в какой папке содержится диплом.

EXAMPLE
input
2
2 1
output
1
В примере Иван может просмотреть папку 2 за 1 секунду и, не найдя там диплома, понять, что диплом находится в папке 1.
Если же он найдет диплом в папке 2, то на поиск уйдет также 1 секунда.
*/

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
