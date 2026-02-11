/*
	Что выведет программа?

	Объяснить порядок выполнения defer функций и итоговый вывод.
*/

/*
	package main

	import "fmt"

	func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
	}

	func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
	}

	func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
	}
*/
/*
	Порядок вызова defer функций
	defer функций вызывается в обратном порядке относительно вызова. В каждой функции сначало выполнятся defer, а уже потом управление будет передано main.

	Что выведет программа?
	В функции test() будет вывод 2, т.к возращается именнованная переменная, return ждет выполнения defer, который увеличивает x на 1
	В функции anotherTest() будет вывод 1, т.к return захватывает переменную x и уже после этого идет выполнение defer, которые уже не влияют на значение переменной.
*/

package main

import "fmt"

func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}

func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}

func main() {
	fmt.Println(test())        // 2
	fmt.Println(anotherTest()) // 1
}
