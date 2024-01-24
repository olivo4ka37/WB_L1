package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Примеры различных переменных
	var structVariable struct{}
	var integerVariable int = 42
	var floatVariable float64 = 3.14
	var stringVariable string = "Hello, Golang!"
	var boolVariable bool = true

	// Определение типа переменной
	detectType(structVariable)
	detectType(integerVariable)
	detectType(floatVariable)
	detectType(stringVariable)
	detectType(boolVariable)
}

// Функция для определения типа переменной
func detectType(variable interface{}) {
	// Используем reflect.TypeOf для определения типа переменной
	variableType := reflect.TypeOf(variable)
	fmt.Printf("Тип переменной: %v\n", variableType)
}
