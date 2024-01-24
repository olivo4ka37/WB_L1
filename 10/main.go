package main

import (
	"fmt"
)

func main() {
	// Инициализируем массив со значениями температурных коллебаний
	xArray := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	// Создаем множество int, в значениях которого подмножества float64
	setOfSets := make(map[int]map[float64]struct{})

	// Записываем каждое значение из массива в подмножества float64
	for _, x := range xArray {
		addElement(setOfSets, int(x/10)*10, x) //int(x/10)*10 - Позволяет градировать подмножества в группы с шагом 10.
	}

	// Выводим содержимое множества множеств
	fmt.Println(setOfSets)
}

// Функция для добавления элемента в множество множеств
func addElement(setOfSets map[int]map[float64]struct{}, key int, value float64) {
	// Если множество для данного ключа еще не существует, создаем его
	if _, ok := setOfSets[key]; !ok {
		setOfSets[key] = make(map[float64]struct{})
	}

	// Добавляем элемент в множество
	setOfSets[key][value] = struct{}{}
}
