/*
	Разработать программу, которая переворачивает порядок слов в строке.
	Пример: входная строка:
	«snow dog sun», выход: «sun dog snow».
	Считайте, что слова разделяются одиночным пробелом.
	Постарайтесь не использовать дополнительные срезы, а выполнять операцию «на месте».
*/

package main

import "fmt"

func inversStr(str string) string {
	sliceRune := []rune(str)
	for i := 0; i < len(sliceRune) / 2; i++ {
		sliceRune[i], sliceRune[len(sliceRune)-1 -i] = sliceRune[len(sliceRune)-1 -i], sliceRune[i]
	}

	startWord := 0
	for endWord := range sliceRune {
		if sliceRune[endWord] == ' ' || endWord == len(sliceRune) -1 {
			l := startWord
			r := endWord
		
			if sliceRune[endWord] == ' ' {
				r = endWord - 1
			}
		
			for l < r {
				sliceRune[l], sliceRune[r] = sliceRune[r], sliceRune[l]
				l++
				r--
			}
			startWord = endWord + 1
		}
	}

	return string(sliceRune)
}

func main() {
	str := "snow dog sun"
	fmt.Println(inversStr(str))
}
