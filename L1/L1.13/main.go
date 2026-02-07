/*
	Обмен значениями без третьей переменной

	Поменять местами два числа без использования временной переменной.

	Подсказка: примените сложение/вычитание или XOR-обмен.
*/

package main

import "fmt"

func main() {
	a := 3
	b := 7
	fmt.Println("Начальные значения:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	a = a + b
	b = a - b
	a = a - b
	fmt.Println("Перестановка с помощью сложения и вычитания:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println("Перестановка с помощью XOR:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
