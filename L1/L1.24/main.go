/*
	Расстояние между точками
	Разработать программу нахождения расстояния между двумя точками на плоскости.
	Точки представлены в виде структуры Point с инкапсулированными (приватными) полями x, y (типа float64) и конструктором.
	Расстояние рассчитывается по формуле между координатами двух точек.

	Подсказка: используйте функцию-конструктор NewPoint(x, y), Point и метод Distance(other Point) float64.
*/

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (this *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow((this.x - other.x), 2) + math.Pow((this.y - other.y), 2))
}

func main() {
	point1 := NewPoint(1, 3)
	point2 := NewPoint(5, 1)
	fmt.Println(point1.Distance(point2))
}