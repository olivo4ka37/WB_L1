package main

import (
	"fmt"
	"strings"
)

func main() {
	// Создаем переменную h типа Human
	// присваеваем значение h.name = "Heisenberg"
	// присваеваем значение h.profession = "School chemistry teacher"
	h := Human{
		name:       "Heisenberg",
		profession: "School chemistry teacher",
	}
	// Создаем переменную a типа Action, эта перменная наследует поля и методы структуры Human
	a := Action{h}
	// Пишем значение вызываемого метода из структуры Human (IntoduceYourSelf)
	fmt.Println(a.actingPerson.IntoduceYourSelf())
	// Пишем значение вызываемого метода из структуры Human (SayMyName)
	fmt.Println(a.actingPerson.SayMyName())
	// Пишем значение вызываемого метода из структуры Action, данный метод использует поля структуры Human (IntoduceYourSelfLouder)
	fmt.Println(a.IntoduceYourSelfLouder())
}

// Human структура содержащая в себе два поля "profession" и "name"
type Human struct {
	profession string
	name       string
}

// SayMyName метод структуры Human, возвращающий "Say my name"
func (h *Human) SayMyName() string {
	return "Say my name"
}

// IntoduceYourSelf метод структуры Human, возвращающий строку с именем и профессей
func (h *Human) IntoduceYourSelf() string {
	return "Hello my name is: " + h.name + " I'm " + h.profession
}

// Action структура, наследующая поля и методы от структуры Human
type Action struct {
	actingPerson Human
}

// IntoduceYourSelfLouder метод стуктуры Action
func (a *Action) IntoduceYourSelfLouder() string {
	return "HELLO MY NAME IS: " + strings.ToUpper(a.actingPerson.name) + " I'M " + strings.ToUpper(a.actingPerson.profession)
}
