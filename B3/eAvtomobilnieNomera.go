package main

import "fmt"

/*
Неизвестный водитель совершил ДТП и скрылся с места происшествия.
Полиция опрашивает свидетелей. Каждый из них говорит, что
запомнил какие-то буквы и цифры номера. Но при этом свидетели не помнят
порядок этих цифр и букв. Полиция хочет проверить несколько подозреваемых автомобилей.
Будем говорить, что номер согласуется с показанием свидетеля, если все символы, которые назвал
свидетель, присутствуют в этом номере (не важно, сколько раз).

INPUT
Сначала задано число  - количество свидетелей.
Далее идет M строк, каждая из которых описывает показания очередного свидетеля.
Эти строки непустые и состоят из не более чем 20 символов.
Каждый символ в строке - либо цифра, либо заглавная латинская буква, причём символы могут повторяться.
Затем идёт число  - количество номеров. Следующие строки представляют из себя номера подозреваемых машин
и имеют такой же формат, как и показания свидетелей.

OUTPUT
Выпишите номера автомобилей, согласующиеся с максимальным количеством свидетелей.
Если таких номеров несколько, то выведите их в том же порядке, в котором они были заданы на входе.

# EXAMPLE 1

input
3
ABC
A37
BCDA
2
A317BD
B137AC

output
B137AC

# EXAMPLE 2

input
2
1ABC
3A4B
3
A143BC
C143AB
AAABC1

output
A143BC
C143AB
*/
type Set struct {
	array            [][]uint8
	numberOfElements uint8
	sizeOf           uint8
}

func hash(a uint8, sizeOf uint8) uint8 {
	return a % sizeOf
}

func (s *Set) find(element uint8) bool {
	for _, now := range s.array[hash(element, s.sizeOf)] {
		if element == now {
			return true
		}
	}
	return false
}
func (s *Set) add(newElement uint8) {
	s.array[hash(newElement, s.sizeOf)] = append(s.array[hash(newElement, s.sizeOf)], newElement)
	s.numberOfElements++
}
func main() {
	var nOfEvidence int
	fmt.Scan(&nOfEvidence)
	Evidences := make([]string, nOfEvidence)
	var currentEvidence string
	for i := 0; i < nOfEvidence; i++ {
		fmt.Scan(&currentEvidence)
		Evidences[i] = currentEvidence
	}
	var nOfSuspects, currentConsistent int
	fmt.Scan(&nOfSuspects)
	var SuspectSet Set
	var Suspect string
	var mostSuspects []string
	var isConsistent bool

	maxCoincide := 0
	for i := 0; i < nOfSuspects; i++ {
		fmt.Scan(&Suspect)
		SuspectSet = Set{
			array:            make([][]uint8, 10),
			sizeOf:           10,
			numberOfElements: 0,
		}
		for k := 0; k < len(Suspect); k++ {
			SuspectSet.add(Suspect[k])
		}
		currentConsistent = 0
		for j := 0; j < nOfEvidence; j++ {
			isConsistent = true
			for k := 0; k < len(Evidences[j]); k++ {
				if !SuspectSet.find(Evidences[j][k]) {
					isConsistent = false
				}
			}
			if isConsistent {
				currentConsistent++
			}
		}
		if currentConsistent > maxCoincide {
			maxCoincide = currentConsistent
			mostSuspects = []string{Suspect}
		} else if currentConsistent == maxCoincide {
			mostSuspects = append(mostSuspects, Suspect)
		}
	}
	for _, c := range mostSuspects {
		fmt.Println(c)
	}
}
