package main

import (
	"fmt"
	"math"
)

// Points - Структура имитирующая координаты x и y
type Points struct {
	x float64
	y float64
}

// New - Конструктор структуры Points
func New(x, y float64) Points {
	return Points{
		x: x,
		y: y,
	}
}

// CalculateDistance - Функция расчета дистанции между двумя точками по координатам x и y
func CalculateDistance(p1 Points, p2 Points) float64 {
	result := math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
	return result
}

func main() {
	// Инициализируем две точки
	point1 := New(1, 1)
	point2 := New(2, 2)

	// Проверяем работу функции CalculateDistance
	fmt.Println(CalculateDistance(point1, point2))

}
