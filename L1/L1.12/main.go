// Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree").
// Создать для неё собственное множество.
// Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.

package main

import "fmt"

func main() {
	array1 := []string{"cat", "cat", "dog", "cat", "tree"}
	array2 := []string{}

	for i := range array1 {
		flag := false
		for j := range array2 {
			if array1[i] == array2[j] {
				flag = true
				break
			}
		}

		if !flag {
			array2 = append(array2, array1[i])
		}
	}

	fmt.Println(array1)
	fmt.Println(array2)
}
