// Реализовать постоянную запись данных в канал (в главной горутине).
// Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.
// Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров.

package main

import (
	"fmt"
	"time"
)

func worker(numbersChanel chan int, id int) {
	for number := range numbersChanel {
		fmt.Println("Worker", id, ":", number)
		time.Sleep(1 * time.Second) // для читаемого вывода
	}
}

func main() {
	var countWorkers int
	fmt.Scan(&countWorkers)

	numbersChanel := make(chan int, countWorkers)

	for i := 0; i < countWorkers; i++ {
		go worker(numbersChanel, i+1)
	}

	for {
		numbersChanel <- 10
	}
}
