/*
	Уникальные символы в строке

	Разработать программу, которая проверяет,
	что все символы в строке встречаются один раз (т.е. строка состоит из уникальных символов).

	Вывод: true, если все символы уникальны, false, если есть повторения.
	Проверка должна быть регистронезависимой, т.е. символы в разных регистрах считать одинаковыми.

	Например: "abcd" -> true, "abCdefAaf" -> false (повторяются a/A), "aabcd" -> false.

	Подумайте, какой структурой данных удобно воспользоваться для проверки условия.
*/

package main

import (
	"fmt"
	"strings"
)

func findElem(str string) bool {
	sliceRune := []rune((strings.ToLower(str)))
	resultMap := make(map[rune]bool)
	for _, v := range sliceRune {
		if resultMap[v] {
			return false
		}
		resultMap[v] = true
	}

	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAa"
	str3 := "aabcd"

	fmt.Println(findElem(str1))
	fmt.Println(findElem(str2))
	fmt.Println(findElem(str3))
}
