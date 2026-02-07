/*
	Паттерн «Адаптер»

	Реализовать паттерн проектирования «Адаптер» на любом примере.

	Описание: паттерн Adapter позволяет сконвертировать интерфейс одного класса в интерфейс другого,
	который ожидает клиент.

	Продемонстрируйте на простом примере в Go: у вас есть существующий интерфейс (или структура) и другой,
	несовместимый по интерфейсу потребитель — напишите адаптер,
	который реализует нужный интерфейс и делегирует вызовы к встроенному объекту.

	Поясните применимость паттерна, его плюсы и минусы, а также приведите реальные примеры использования.
*/

package main

import "fmt"

// уже написанный код, который нельзя менять
type LegacyLogger struct{}

func (l *LegacyLogger) WriteMessage(msg string) {
	fmt.Println("Legacy log:", msg)
}

// клиент
type Logger interface {
	Log(message string)
}

// адаптер
type LoggerAdapter struct {
	legacy *LegacyLogger
}

func (a *LoggerAdapter) Log(message string) {
	a.legacy.WriteMessage(message)
}

func main() {
	legacy := &LegacyLogger{}
	adapter := &LoggerAdapter{legacy: legacy}

	var logger Logger = adapter
	logger.Log("Hello Adapter pattern")
}
