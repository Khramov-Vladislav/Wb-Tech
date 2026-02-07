/*
	Определение типа переменной в runtime

	Разработать программу, которая в runtime способна определить тип переменной,
	переданной в неё (на вход подаётся interface{}).
	Типы, которые нужно распознавать: int, string, bool, chan (канал).

	Подсказка: оператор типа switch v.(type) поможет в решении.
*/

package main

import "fmt"

func checkType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Printf("тип %T\n", v)
	case string:
		fmt.Printf("тип %T\n", v)
	case bool:
		fmt.Printf("тип %T\n", v)
	case chan int:
		fmt.Printf("тип %T\n", v)
	case chan string:
		fmt.Printf("тип %T\n", v)
	case chan bool:
		fmt.Printf("тип %T\n", v)
	default:
		fmt.Println("Неизвестный тип данных")
	}
}

func main() {
	var typeInt int = 10
	var typeString string = "Hello"
	var typeBool bool = true
	var typeChanInt chan int
	var typeChanString chan string
	var typeChanBool chan bool

	checkType(typeInt)
	checkType(typeString)
	checkType(typeBool)
	checkType(typeChanInt)
	checkType(typeChanString)
	checkType(typeChanBool)
}
