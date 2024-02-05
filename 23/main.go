package main

import (
	"fmt"
)

func main() {
	// Инициализация слайса
	xArray := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Проверка работы программы
	fmt.Println(removElement(xArray, 9))
	fmt.Println(removElement(xArray, 99))
}

// removElement Удаляет i-ый элемент слайса
// Если введен пустой массив или позиция i выходит за предел слайса
// функция выведет сообщение и вернет нетронутый слайс.
func removElement(xArray []int, i int) []int {
	if len(xArray) == 0 {
		fmt.Println("You can't use a zero len slice to this function")
		return xArray
	}

	if i >= len(xArray) {
		fmt.Println("Index i out of range of slice")
		return xArray
	}

	return append(xArray[:i], xArray[i+1:]...)
}
