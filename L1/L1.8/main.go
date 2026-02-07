/*
	Установка бита в числе

	Дана переменная типа int64. Разработать программу, которая устанавливает i-й бит этого числа в 1 или 0.

	Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).

	Подсказка: используйте битовые операции (|, &^).
*/

package main

import (
	"fmt"
)

func main() {
	var number int64 = 5
	fmt.Printf("number: %d = %b\n", number, number)
	var index int = 0

	var mask int64 = 1 << index
	fmt.Printf("mask: %d = %b\n", mask, mask)

	if mask&number != 0 {
		fmt.Printf("%d = %b\n", number&^mask, number&^mask)
	} else {
		fmt.Printf("%d = %b\n", number|mask, number|mask)
	}
}
