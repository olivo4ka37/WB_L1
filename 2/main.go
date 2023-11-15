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

	// Добавляем в WaitGroup количество ожидаемых горутин равное кол-ву элементов массива nums
	wg.Add(len(nums))

	// Создаем цикл операций для каждого элемента массива
	for i := 0; i < len(nums); i++ {
		// Вызываем горутину считающую квадрат nums[i]
		// Выводящую результат в stdout
		go func(i int) {
			squareOfInt(nums[i], &wg)
		}(i)
	}

	// Wait() Блокируется пока счетчик не станет равным нулю
	// Это позволяет дождаться выполнения всех горутин, а только потом завершить программу.
	wg.Wait()
}

// squareOfInt Функция считающая и выводящая квадрат числа, при каждом выполнении
// уменьшает счетчик операций WaitGroup на 1.
func squareOfInt(i int, wg *sync.WaitGroup) {
	fmt.Printf("square of %d is equal to %d\n", i, i*i)
	wg.Done()
}
