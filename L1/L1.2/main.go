/*
	Конкурентное возведение в квадрат

	Написать программу, которая конкурентно рассчитает значения квадратов чисел, взятых из массива [2,4,6,8,10],
	и выведет результаты в stdout.

	Подсказка: запусти несколько горутин, каждая из которых возводит число в квадрат.
*/

package main

import (
	"fmt"
	"sync"
)

func Squaring(wg *sync.WaitGroup, number int, chanel chan int) {
	defer wg.Done()
	chanel <- number * number
}

func main() {
	var wg sync.WaitGroup

	array := [5]int{2, 4, 6, 8, 10}
	chanel := make(chan int, len(array))

	for _, v := range array {
		wg.Add(1)
		go Squaring(&wg, v, chanel)
	}
	wg.Wait()
	close(chanel)

	for v := range chanel {
		fmt.Println(v)
	}
}
