package main

import (
	"fmt"
	"time"
)

func main() {
	// Проверка работы первой функции
	sleep1(3 * time.Second)
	fmt.Println("TEST MEOW")

	// Проверка работы второй функции
	sleep2(3 * time.Second)
	fmt.Println("TEST MEOW")
}

// sleep1 Ждёт указанное время t, потом возвращает текущее время по возвращаемуму каналу
func sleep1(t time.Duration) {
	<-time.After(t)
}

// sleep2 Создаем новый таймер, который возвращает текущее время по каналу не раньше чем через время t
// Канал принадлежит структуре Timer тип канала time
func sleep2(t time.Duration) {
	timer := time.NewTimer(t)

	<-timer.C
}
