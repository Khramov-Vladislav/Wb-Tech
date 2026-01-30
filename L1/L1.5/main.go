package main

import (
	"fmt"
	"time"
)

func main() {
	N := 5 * time.Second // сколько секунд работает программа
	ch := make(chan int)

	// таймер завершения
	timeout := time.After(N)

	// Отправитель
	go func() {
		i := 1
		for {
			select {
			case ch <- i:
				i++
				time.Sleep(500 * time.Millisecond) // чтобы не спамило слишком быстро
			case <-timeout:
				fmt.Println("Отправитель остановлен")
				return
			}
		}
	}()

	// Приёмник
	for {
		select {
		case val := <-ch:
			fmt.Println("Получено:", val)
		case <-timeout:
			fmt.Println("Приёмник остановлен")
			fmt.Println("Программа завершена")
			return
		}
	}
}
