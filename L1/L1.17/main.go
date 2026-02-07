/*
	Бинарный поиск

	Реализовать алгоритм бинарного поиска встроенными методами языка.
	Функция должна принимать отсортированный слайс и искомый элемент,
	возвращать индекс элемента или -1, если элемент не найден.

	Подсказка: можно реализовать рекурсивно или итеративно, используя цикл for.
*/

package main

import "fmt"

func binarySearch(array []int, findNumber int) int {
	left := 0
	right := len(array) - 1

	for right-left > 1 {
		middle := (left + right) / 2
		if array[middle] >= findNumber {
			right = middle
		} else {
			left = middle
		}
	}

	if array[right] == findNumber {
		return right
	} else {
		return -1
	}
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	findNumber := 8
	fmt.Println(binarySearch(array, findNumber))
}
