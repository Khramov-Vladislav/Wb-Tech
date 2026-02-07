/*
	Остановка горутины

	Реализовать все возможные способы остановки выполнения горутины.

	Классические подходы: выход по условию, через канал уведомления, через контекст, прекращение работы runtime.Goexit() и др.

	Продемонстрируйте каждый способ в отдельном фрагменте кода.
*/

package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// выход по условию
	wg.Go(func() {
		fmt.Println("Горутина с выходом по условию начала работу")
		for n := range 10 {
			fmt.Println("Горутина 1:", n)
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println("Горутина с выходом по условию завершила работу")
	})

	// выход через канал уведомления TODO
	stopChanel := make(chan bool)
	wg.Go(func() {
		fmt.Println("Горутина с выходом через канал уведомления начала работу")
		n := 0
		for {
			select {
			case <-stopChanel:
				fmt.Println("Горутина с выходом через канал уведомления закончила работу")
				return
			default:
				fmt.Println("Горутина 2:", n)
				n++
				time.Sleep(100 * time.Millisecond)
			}
		}
	})

	wg.Go(func() {
		stop := time.After(1 * time.Second)
		<-stop
		stopChanel <- true
	})

	// выход через закрытие канала
	closeChanel := make(chan int)
	wg.Go(func() {
		fmt.Println("Горутина с выходом через закрытие канала начала работу")
		n := 0
		for {
			select {
			case _, ok := <-closeChanel:
				if !ok {
					fmt.Println("Горутина с выходом через закрытие канала завершила работу")
					return
				}
			default:
				fmt.Println("Горутина 3:", n)
				n++
				time.Sleep(100 * time.Millisecond)
			}
		}
	})

	wg.Go(func() {
		stop := time.After(1 * time.Second)
		<-stop
		close(closeChanel)
	})

	// выход через context
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	wg.Go(func() {
		fmt.Println("Горутина с выходом через context начала работу")
		n := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина с выходом через context завершила работу")
				return
			default:
				fmt.Println("Горутина 4:", n)
				n++
				time.Sleep(100 * time.Millisecond)
			}
		}
	})

	// прекращение работы runtime.Goexit()
	wg.Go(func() {
		fmt.Println("Горутина с прекращением работы runtime.Goexit() начала работу")
		defer fmt.Println("Горутина с прекращением работы runtime.Goexit() завершила работу")

		n := 0
		for {
			if n == 10 {
				runtime.Goexit()
			}
			fmt.Println("Горутина 5:", n)
			time.Sleep(100 * time.Millisecond)
			n++
		}
	})

	wg.Wait()
	fmt.Println("main завершила работу")
}
