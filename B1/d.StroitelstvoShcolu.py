# STATEMENT
# В деревне Интернетовка все дома расположены вдоль одной улицы по одну сторону от нее.
# По другую сторону от этой улицы пока ничего нет, но скоро все будет – школы, магазины, кинотеатры и т.д.
# Для начала в этой деревне решили построить школу. Место для строительства школы решили выбрать так,
# чтобы суммарное расстояние, которое проезжают ученики от своих домов до школы, было минимально.
# План деревни можно представить в виде прямой, в некоторых целочисленных точках которой находятся дома учеников.
# Школу также разрешается строить только в целочисленной точке этой прямой (в том числе разрешается строить школу в
# точке, где расположен один из домов – ведь школа будет расположена с другой стороны улицы).
# Напишите программу, которая по известным координатам домов учеников поможет определить координаты места строительства школы.

# INPUT
# Сначала вводится число N — количество учеников (0 < N < 100001). Далее идут в строго возрастающем порядке координаты
# домов учеников — целые числа, не превосходящие 2 × 10^9 по модулю.
# OUTPUT
# Выведите одно целое число — координату точки, в которой лучше всего построить школу. Если ответов несколько, выведите любой из них.
# EXAMPLE 1
# input
# 4
# 1 2 3 4
# output
# 3
# EXAMPLE 2
# input
# 3
# -1 0 1
# output
# 0

N = int(input())
b = [int(i) for i in input().split()]
print(b[int(N/2)])