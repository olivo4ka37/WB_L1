package main

import (
	"fmt"
)

func setBit(value int64, pos uint, bitValue uint) int64 {
	// Очищаем i-й бит
	mask := ^(int64(1) << pos)
	x1 := fmt.Sprintf("%b", mask)
	fmt.Println(x1)
	x1 = fmt.Sprintf("%b", value)
	fmt.Println(x1)
	value &= mask
	x1 = fmt.Sprintf("%b", value)
	fmt.Println("Valeus &= mask is: ", x1)

	// Устанавливаем i-й бит в новое значение
	value |= int64(bitValue) << pos
	x1 = fmt.Sprintf("%b", value)
	fmt.Println(x1)

	return value
}

func main() {
	var number int64 = 7
	var bitPosition uint = 2
	var bitValue uint = 1

	// Устанавливаем i-й бит в 1
	result := setBit(number, bitPosition, bitValue)

	fmt.Printf("Исходное значение: %d\n", number)
	fmt.Printf("После установки i-го бита: %d\n", result)
}
