package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x uint64

	x = 123456789123456789123456789123456789123456789123456789123456789123456789

	fmt.Println(x)

}

func bigMath(a int, b int, operand byte) string {

	return ""
}

// Конвертирует десятичное число в двоичную и возвращает значение строкой
func ConvertTo2Base(x int) (string, error) {
	str := strconv.Itoa(x)
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 2), nil
}
