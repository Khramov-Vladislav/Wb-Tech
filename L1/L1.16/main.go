/*
	Быстрая сортировка (quicksort)

	Реализовать алгоритм быстрой сортировки массива встроенными средствами языка.
	Можно использовать рекурсию.

	Подсказка: напишите функцию quickSort([]int) []int которая сортирует срез целых чисел.
	Для выбора опорного элемента можно взять середину или первый элемент.
*/

package main

import (
	"fmt"
)

func quickSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	mainElement := array[len(array)/2]
	var leftPart []int
	var rightPart []int
	var medium []int

	for i := range array {
		if array[i] < mainElement {
			leftPart = append(leftPart, array[i])
		} else if array[i] == mainElement {
			medium = append(medium, array[i])
		} else if array[i] > mainElement {
			rightPart = append(rightPart, array[i])
		}
	}

	sortedLeftPart := quickSort(leftPart)
	sortedRightPart := quickSort(rightPart)

	var result []int
	result = append(result, sortedLeftPart...)
	result = append(result, medium...)
	result = append(result, sortedRightPart...)

	return result
}

func main() {
	array := []int{1, 2, 10, 6, 3, 8, 0, 4, 2, 7}
	fmt.Println(quickSort(array))
}
