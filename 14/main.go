package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Примеры различных переменных
	var integerVariable int = 42
	var stringVariable string = "Hello, Golang!"
	var boolVariable bool = true
	var chanVariable chan struct{}

	// Определение типа переменной
	detectType(integerVariable)
	detectType(stringVariable)
	detectType(boolVariable)
	detectType(chanVariable)
}

// Функция для определения типа переменной
func detectType(variable interface{}) {
	// Используем reflect.TypeOf для определения типа переменной
	variableType := reflect.TypeOf(variable)
	fmt.Printf("Тип переменной: %v\n", variableType)
}
