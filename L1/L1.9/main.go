// Разработать конвейер чисел. Даны два канала: в первый пишутся числа x из массива, во второй – результат операции x*2.
// После этого данные из второго канала должны выводиться в stdout.
// То есть, организуйте конвейер из двух этапов с горутинами: генерация чисел и их обработка.
// Убедитесь, что чтение из второго канала корректно завершается.

package main

import (
	"fmt"
	"sync"
)

func writeFromArray(numbersChanel chan<- int, array []int) {
	fmt.Println("Горутина записи из массива начала работу")
	for i := range array {
		numbersChanel <- array[i]
		fmt.Println("Отправлено значение:", array[i])
	}
	close(numbersChanel)
	fmt.Println("Горутина записи из массива закончила работу")
}

func multiolecateNumbersFromChanel(numbersChanel <-chan int, multiplecateChanel chan<- int) {
	fmt.Println("Горутина умножения числа начала работу")
	for {
		val, ok := <-numbersChanel
		if !ok {
			fmt.Println("Горутина умножения числа закончила работу")
			close(multiplecateChanel)
			return
		}
		multiplecateChanel <- val * 2
		fmt.Println("Произведение:", val*2)
	}
}

func readNumbersFromChanel(multiplecateChanel <-chan int) {
	fmt.Println("Горутина чтения из канала начала работу")
	for {
		val, ok := <-multiplecateChanel
		if !ok {
			fmt.Println("Горутина чтения из канала закончила работу")
			return
		}

		fmt.Println("Получено значение:", val)
	}
}

func main() {
	var wg sync.WaitGroup
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	numbersChanel := make(chan int)
	multiplecateChanel := make(chan int)

	wg.Go(func() {
		writeFromArray(numbersChanel, array[:])
	})

	wg.Go(func() {
		multiolecateNumbersFromChanel(numbersChanel, multiplecateChanel)
	})

	wg.Go(func() {
		readNumbersFromChanel(multiplecateChanel)
	})

	wg.Wait()
	fmt.Println("main завершала работу")
}
