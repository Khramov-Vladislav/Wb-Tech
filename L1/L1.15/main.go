/*
	Небольшой фрагмент кода — проблемы и решение

	Рассмотреть следующий код и ответить на вопросы:
	к каким негативным последствиям он может привести и как это исправить?

	Приведите корректный пример реализации.
	Вопрос: что происходит с переменной justString?
*/

/*
	var justString string

	func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
	}

	func main() {
	someFunc()
	}
*/

/*
	проблема в том, что нам нужна строка в 100 символов, а будет храниться вся исходная строка на 1024 сивола,
	что не эфективно
*/

package main

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	newString := v[:100]
	justString = string(newString)
}

func main() {
	someFunc()
}
