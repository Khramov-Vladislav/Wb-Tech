/*
	Своя функция Sleep
	Реализовать собственную функцию sleep(duration) аналогично встроенной функции time.Sleep,
	которая приостанавливает выполнение текущей горутины.

	Важно: в отличие от настоящей time.Sleep,
	ваша функция должна именно блокировать выполнение (например, через таймер или цикл),
	а не просто вызывать time.Sleep :) — это упражнение.

	Можно использовать канал + горутину, или цикл на проверку времени (не лучший способ, но для обучения).
*/

package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	done := make(chan struct{})

	go func() {
		start := time.Now()
		for time.Since(start) < duration {
		}

		close(done)
	}()

	<-done
}

func main() {
	fmt.Println("main начала работу")
	sleep(2 * time.Second)
	fmt.Println("закончила работу")
}
