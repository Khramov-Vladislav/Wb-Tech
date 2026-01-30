// Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала – читать эти значения.
// По истечении N секунд программа должна завершаться.
// Подсказка: используйте time.After или таймер для ограничения времени работы.

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var N time.Duration
	fmt.Scan(&N)

	numbersChanel := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// запись
	wg.Add(1)
	go func() {
		number := 0
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				fmt.Println("Горутина для записи завершила работу.")
				return
			case numbersChanel <- number:
				time.Sleep(100 * time.Millisecond) // для читаемого вывода
			}
			number++
		}
	}()

	// чтение
	wg.Add(1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				fmt.Println("Горутина для чтения завершила работу.")
				return
			case number := <-numbersChanel:
				fmt.Println("Number:", number)
			}
		}
	}()

	// остановка
	timeChanel := time.After(N * time.Second)
	<-timeChanel
	cancel()
	wg.Wait()
	fmt.Println("Программа завершилась")
}
