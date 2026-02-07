/*
	Конкурентный счетчик

	Реализовать структуру-счётчик, которая будет инкрементироваться в конкурентной среде (т.е. из нескольких горутин).
	По завершению программы структура должна выводить итоговое значение счётчика.

	Подсказка: вам понадобится механизм синхронизации, например, sync.Mutex или sync/Atomic для безопасного инкремента.
*/

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	number int
	mtx    sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{}
}

func count(C *Counter) {
	C.mtx.Lock()
	C.number++
	C.mtx.Unlock()
}

func main() {
	var wg sync.WaitGroup

	c := NewCounter()
	for range 10000 {
		wg.Go(func() {
			count(c)
		})
	}

	wg.Wait()
	fmt.Println(c.number)
}
