/*
	Конкурентная запись в map

	Реализовать безопасную для конкуренции запись данных в структуру map.

	Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).

	Проверьте работу кода на гонки (util go run -race).
*/

package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, lock *sync.Mutex, id int, numberMap map[int]int) {
	fmt.Println("Горутина", id, "начала работу")

	lock.Lock()
	numberMap[id] = id * 10
	lock.Unlock()

	fmt.Println("Горутина", id, "закончила работу")
	wg.Done()
}

func main() {
	numberMap := make(map[int]int)

	var wg sync.WaitGroup
	var lock sync.Mutex

	countGorutins := 10

	for i := range countGorutins {
		wg.Add(1)
		go worker(&wg, &lock, i, numberMap)
	}

	wg.Wait()

	fmt.Println("Длинна map равна", len(numberMap))
	fmt.Println("main завершила работу")
}
