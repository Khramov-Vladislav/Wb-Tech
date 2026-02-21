/*
	Получение точного времени (NTP).

	Создать программу, печатающую точное текущее время с использованием NTP-сервера.

	* Реализовать проект как модуль Go.

	* Использовать библиотеку ntp для получения времени.

	* Программа должна выводить текущее время, полученное через NTP (Network Time Protocol).

	* Необходимо обрабатывать ошибки библиотеки: в случае ошибки вывести её текст в STDERR и вернуть ненулевой код выхода.

	* Код должен проходить проверки (vet и golint), т.е. быть написан идиоматически корректно.
*/

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func getNTP() (time.Time, error) {
	ntpServer := "pool.ntp.org"
	timeNTP, err := ntp.Time(ntpServer)
	return timeNTP, err
}

func main() {
	timeNTP, err := getNTP()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка получения времени NTP", err)
		os.Exit(1)
	}

	fmt.Println("Точное время (NTP):", timeNTP)
}
