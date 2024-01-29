package main

import (
	"fmt"
	"time"
)

// У нас была структура InputSystem в который был обычный спаммер
// через эту структуру потребовалось использовать новые методы спамма которые
// находились в AdvancedInputSystem (ButtonSpammer, MouseClicker).
//
// Соответсвенно были интерфейсы DefaultSpammer и AdvancedSpammer.
// Для выполнения поставленной задачи была создана структура  SpamAdapter в которой
// мы реализовали интерфейс AdvancedInputSystem в переменной advancedcComputerSpammer
// а в структуру InputSystem добавляем перменную spamAdapter типа SpamAdapter
// Тем самым мы теперь можем использовать методы AdvancedSpammer в структуре InputSystem.
func main() {
	s := InputSystem{}

	s.Spammer(3, 100)
	s.spamAdapter.advancedcComputerSpammer.ButtonSpammer(3, 100, "F")
	s.spamAdapter.advancedcComputerSpammer.MouseClicker(3, 100, "LeftClick")
}

// DefaultSpammer интерфейс обычного спаммера
type DefaultSpammer interface {
	Spammer(count int, delay time.Duration) error
}

// AdvancedSpammer интерфейс продвинутого спаммера
type AdvancedSpammer interface {
	ButtonSpammer(count int, delay time.Duration, button string) error
	MouseClicker(count int, delay time.Duration, button string) error
}

// InputSystem Структура в которую нам нужно было добавить методы продвинутого спаммера
type InputSystem struct {
	spamAdapter SpamAdapter
}

// AdvancedInputSystem Структура продвинутого спаммера
type AdvancedInputSystem struct{}

// ButtonSpammer Спаммер отдельных кнопок клавиатуры функция структуры AdvancedInputSystem
func (AdvancedInputSystem) ButtonSpammer(count int, delay time.Duration, button string) error {
	for i := 0; i < count; i++ {
		fmt.Println("I PRESSED THIS BUTTON: ", button)
		time.Sleep(time.Millisecond * delay)
	}

	return nil
}

// ButtonSpammer Спаммер отдельных кнопок мыши функция структуры AdvancedInputSystem
func (AdvancedInputSystem) MouseClicker(count int, delay time.Duration, button string) error {
	for i := 0; i < count; i++ {
		fmt.Println("I Click This mouse button: ", button)
		time.Sleep(time.Millisecond * delay)
	}

	return nil
}

// SpamAdapter Адаптер структуры AdvancedInputSystem
type SpamAdapter struct {
	advancedcComputerSpammer AdvancedInputSystem
}

// Spammer функция простого спамма структуры InputSystem
func (in InputSystem) Spammer(count int, delay time.Duration) error {
	for i := 0; i < count; i++ {
		fmt.Println("SPAM")
		time.Sleep(time.Millisecond * delay)
	}

	return nil
}
