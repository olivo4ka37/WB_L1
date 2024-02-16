package main

import (
	"fmt"
)

func main() {
	xArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	// Тест кейс на несуществующий элемент, первый элемент списка, последний элемент списка
	fmt.Println(binarySearch(xArray, 0, len(xArray)-1, 0))
	fmt.Println(binarySearch(xArray, 0, len(xArray)-1, 1))
	fmt.Println(binarySearch(xArray, 0, len(xArray)-1, 11))
}

// binarySearch реализация бинарного поиска на go
func binarySearch(arr []int, left, right, searchX int) int {

	// Вычисляем середину заданного отрезка
	middle := (right + left) / 2

	// Поиск устроен рекурсий, суть заключается в деление заданного отрезка массива на две части
	//и сравнивание середины с элементом который мы ищем, если элемент больше ищем в правой части поделенного массива
	//если элемент меньше ищем в левой части поделенного массива.
	if left <= right {

		if arr[middle] == searchX {
			return middle
		} else if arr[middle] < searchX {
			return binarySearch(arr, middle+1, right, searchX)
		} else {
			return binarySearch(arr, left, middle-1, searchX)
		}

	} else { //Если элемента который мы ищем, нету в массиве, тогда одна из границ будет сдвигаться пока не станет больше или меньше другой
		// и тогда программа выведет -1.
		return -1
	}

}
