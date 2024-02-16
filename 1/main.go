package main

import (
	"fmt"
	"strings"
)

func main() {
	// Создаем переменную h типа human
	// присваеваем значение h.name = "Heisenberg"
	// присваеваем значение h.profession = "School chemistry teacher"
	h := human{
		name:       "Heisenberg",
		profession: "School chemistry teacher",
	}
	// Создаем переменную a типа action, эта перменная наследует поля и методы структуры human
	a := action{h}
	// Пишем значение вызываемого метода из структуры human (IntoduceYourSelf)
	fmt.Println(a.actingPerson.IntoduceYourSelf())
	// Пишем значение вызываемого метода из структуры human (SayMyName)
	fmt.Println(a.actingPerson.SayMyName())
	// Пишем значение вызываемого метода из структуры action, данный метод использует поля структуры human (IntoduceYourSelfLouder)
	fmt.Println(a.IntoduceYourSelfLouder())
}

// human структура содержащая в себе два поля "profession" и "name"
type human struct {
	profession string
	name       string
}

// SayMyName метод структуры human, возвращающий "Say my name"
func (h *human) SayMyName() string {
	return "Say my name"
}

// IntoduceYourSelf метод структуры human, возвращающий строку с именем и профессей
func (h *human) IntoduceYourSelf() string {
	return "Hello my name is: " + h.name + " I'm " + h.profession
}

// action структура, наследующая поля и методы от структуры human
type action struct {
	actingPerson human
}

// IntoduceYourSelfLouder метод стуктуры action
func (a *action) IntoduceYourSelfLouder() string {
	return "HELLO MY NAME IS: " + strings.ToUpper(a.actingPerson.name) + " I'M " + strings.ToUpper(a.actingPerson.profession)

}
