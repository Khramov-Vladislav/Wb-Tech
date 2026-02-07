/*
	Пересечение множеств

	Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов)
	— т.е. вывести элементы, присутствующие и в первом, и во втором.

	Пример:
	A = {1,2,3}
	B = {2,3,4}
	Пересечение = {2,3}
*/

package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := []int{2, 4, 6, 8, 10}
	resultSlice := []int{}

	for i := range slice1 {
		for j := range slice2 {
			if slice1[i] == slice2[j] {
				resultSlice = append(resultSlice, slice1[i])
			}
		}
	}
	fmt.Println(resultSlice)
}
