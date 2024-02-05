package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Инициализация переменных специального типа для работы с большими числами
	x := new(big.Int)
	y := new(big.Int)
	result := new(big.Int)

	// Установка больших значений которые выходят за пределы int64 и uint64
	x.SetString("123456789123456789123456789123456789123456789123456789123456789123456789", 10)
	y.SetString("123456789123456789123456789123456789123456789123456789123456789123456789", 10)

	// Вывод результатов арифимитических операций
	fmt.Printf("a + b = %d\n", result.Add(x, y))
	fmt.Printf("a - b = %d\n", result.Sub(x, y))
	fmt.Printf("a * b = %d\n", result.Mul(x, y))
	fmt.Printf("a / b = %d\n", result.Div(x, y))
}
