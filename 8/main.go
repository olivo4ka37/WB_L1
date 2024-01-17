package main

import (
	"fmt"
)

func setBit(value int64, i uint8, bitValue uint8) int64 {
	// Сдвигаем 1 бит на нужную позицию, создаем обратную маску
	// Используем побитовое сравнение и (and) с этой маской
	// Тем самым очитсили i-ый бит
	value &= ^(int64(1) << i)

	// Создаем маску в которой бит нужной позиции равен указанному пользователем значению
	// Используем побитовое или (or) value с очищенным i-ым битом.
	value |= int64(bitValue) << i

	return value
}

func main() {
	var number int64 = 4
	var bitPosition uint8 = 2
	var bitValue uint8 = 0

	// Устанавливаем i-й бит в 1
	result := setBit(number, bitPosition, bitValue)

	fmt.Printf("Исходное значение: %d\n", number)
	fmt.Printf("После установки i-го бита: %d\n", result)
}
