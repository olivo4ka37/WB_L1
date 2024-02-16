package main

import (
	"fmt"
	"sync"
)

func main() {
	// Инициализируем wg как переменную типа WaitGroup
	// wg будем использовать для ожидания ещё не завершенных операций
	var wg sync.WaitGroup

	// Инициализируем и присваиваем значение массиву nums
	nums := []int{2, 4, 6, 8, 10}
	// Переменная для хранения суммы квадратов массива nums
	result := 0

	// Создаем Канал в котором будем передавать значения квадратов nums[i] между горутинами
	squaresCh := make(chan int, len(nums))

	// Создаем цикл операций для вызова горутин
	for i := 0; i < len(nums); i++ {
		// Увеличиваем счетчик WaitGroup на 1
		wg.Add(1)
		//Вызываем горутину считающую квадрат nums[i] и записывающую её
		// в канал squaresCh.
		go func(i int) {
			addSquareToChanel(nums[i], squaresCh, &wg)
		}(i)
	}

	// Читаем из канала результаты работы горутин и суммируем их в переменной result
	for i := 0; i < cap(squaresCh); i++ {
		result += <-squaresCh
	}

	// Закрываем канал
	close(squaresCh)

	// Выводим результат
	fmt.Printf("Sum squares of array nums is equal to %d", result)

	wg.Wait()
}

// addSquareToChanel Добавляет в канал значение квадрата
func addSquareToChanel(num int, sumCh chan int, wg *sync.WaitGroup) {
	sumCh <- calcSquare(num, wg)
}

// calcSquare Считает квадрат числа и уменьшает счетчик горутин на 1
func calcSquare(num int, wg *sync.WaitGroup) int {
	wg.Done()
	return num * num

}
