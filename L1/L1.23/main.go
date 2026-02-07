/*
	Удаление элемента слайса

	Удалить i-ый элемент из слайса. Продемонстрируйте корректное удаление без утечки памяти.

	Подсказка: можно сдвинуть хвост слайса на место удаляемого элемента (copy(slice[i:], slice[i+1:]))
	и уменьшить длину слайса на 1.
*/

package main

import "fmt"

func delElemSlice[T any](slice []T, idx int) []T {
	copy(slice[idx:], slice[idx+1:])
	newSlice := slice[:len(slice)-1]
	return newSlice
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(delElemSlice(slice, 5))
}
