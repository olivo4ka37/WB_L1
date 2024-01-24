package main

import "fmt"

func main() {
	// Создаем множества int значений
	// x1 и x2 множества между которыми будем искать пересечение
	// result множество в которое будет записанно множество пересечений x1 и x2
	x1 := make(map[int]struct{})
	x2 := make(map[int]struct{})
	result := make(map[int]struct{})

	// Добавляем в множество x1 все значения от 0 до 100
	for i := 0; i < 101; i++ {
		x1[i] = struct{}{}
		// Добавляем в множество x2 все значения которые делятся на 10 без остатка
		if i%10 == 0 {
			x2[i] = struct{}{}
		}
	}

	// Проверяем наличие значений x2 в x1, если есть записываем в result
	for x, _ := range x2 {
		_, contains := x1[x] //contains - булевое значение определяющее наличие данного ключа в множестве
		if contains {
			result[x] = struct{}{}
		}
	}

	fmt.Println(result)
}
