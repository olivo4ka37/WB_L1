package main

import (
	"fmt"
)

func main() {
	xArray := []int{-13, 413, 24, 124, 12, 1, 15632, 234, 85, 753, 9, 13, 3415, 864, 1, 97, 315, 15, -12, 21521}

	quickSort(xArray, 0, len(xArray)-1)

	fmt.Println(xArray)
}

// Реализация быстрой сортировки
func quickSort(lst []int, left int, right int) []int {
	//Создаем копии пришедших переменных, с которыми будем манипулировать в дальнейшем.
	l := left
	r := right

	//Вычисляем 'центр', на который будем опираться. Берем значение ~центральной ячейки массива.
	center := lst[(left+right)/2]

	//Цикл, начинающий саму сортировку
	for l <= r {

		//Ищем значения больше 'центра'
		for lst[r] > center {
			r--
		}

		//Ищем значения меньше 'центра'
		for lst[l] < center {
			l++
		}

		//После прохода циклов проверяем счетчики циклов
		if l <= r {
			//И если условие true, то меняем ячейки друг с другом.
			lst[r], lst[l] = lst[l], lst[r]
			l++
			r--
		}
	}

	if r > left {
		quickSort(lst, left, r)
	}

	if l > right {
		quickSort(lst, l, right)
	}

	return lst

}
