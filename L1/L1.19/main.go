/*
	Разворот строки

	Разработать программу, которая переворачивает подаваемую на вход строку.

	Например: при вводе строки «главрыба» вывод должен быть «абырвалг».

	Учтите, что символы могут быть в Unicode (русские буквы, emoji и пр.),
	то есть просто iterating по байтам может не подойти — нужен срез рун ([]rune).
*/

package main

import (
	"fmt"
)

func inversStr(str string) string {
	sliceRune := []rune(str)
	for i := 0; i < len(sliceRune)/2; i++ {
		sliceRune[i], sliceRune[len(sliceRune)-1-i] = sliceRune[len(sliceRune)-1-i], sliceRune[i]
	}
	return string(sliceRune)
}

func main() {
	str := "главрыба"
	fmt.Println(inversStr(str))
}
