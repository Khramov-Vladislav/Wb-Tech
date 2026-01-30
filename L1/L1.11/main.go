package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств
// (например, двух слайсов) — т.е. вывести элементы,
//  присутствующие и в первом, и во втором.

// Пример:
// A = {1,2,3}
// B = {2,3,4}
// Пересечение = {2,3}

func IntersectionOfSets(setOne, setTwo []int) []int {
	setResult := []int{}
	for i := range setOne {
		for j := range setTwo {
			if setOne[i] == setTwo[j] {
				setResult = append(setResult, setOne[i])
			}
		}
	}

	return setResult
}

// может быть поменять двойной цикл на что-то другое, по другому реализовать

func main() {
	setOne := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	setTwo := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	setResult := IntersectionOfSets(setOne, setTwo)
	fmt.Print(setResult)
}
