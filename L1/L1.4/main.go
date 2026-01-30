// Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).
// Выберите и обоснуйте способ завершения работы всех горутин-воркеров при получении сигнала прерывания.
// Подсказка: можно использовать контекст (context.Context) или канал для оповещения о завершении.

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, numbersChanel chan int, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Завершаю работу", "Worker", id)
			wg.Done()
			return
		case number, ok := <-numbersChanel:
			if !ok {
				wg.Done()
				return
			}
			fmt.Println("Worker", id, ":", number)
		}
	}
}

func main() {
	fmt.Println("Программа запущена, введите кол-во воркеров (Завершение программы: Control+ C )")

	var countWorkers int
	fmt.Scan(&countWorkers)

	numbersChanel := make(chan int, countWorkers)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup

	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel, syscall.SIGINT)

	go func() {
		signal := <-signalChanel
		fmt.Println("\nПолучен сигнал:", signal)
		cancel()
	}()

	for i := 0; i < countWorkers; i++ {
		wg.Add(1)
		go worker(ctx, &wg, numbersChanel, i+1)
	}

	for {
		select {
		case <-ctx.Done():
			close(numbersChanel)
			wg.Wait()
			return
		case numbersChanel <- 10:
			time.Sleep(200 * time.Millisecond) // для читаемого вывода
		}
	}
}

// Обоснование способа завершения работы всех горутин-воркеров при получении сигнала прерывания
// context взят для того чтобы объединить все горутины и одновременно ими управлять, он вызывается если передан сигнал о завешении программы
// также использовал WaightGroup для того чтобы дождаться завершения горутин а уже потом завершшать прошрамму
