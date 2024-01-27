package main

import (
	"fmt"
)

func main() {
	// Инициализируем стэк структуру и тестовую строку
	s := stack{}
	testString := "главрыба_что-то_на_японском_лол)_ ぁ さ ゎ デ ニ ね"

	// Записываем все руны из строки в стэк
	for _, x := range testString {
		s.push(x)
	}

	// Достаем из стэка все что мы туда занесли
	// Тем самым получаем обратную строку
	for range testString {
		fmt.Print(string(s.pop()))
	}

}

// стэк структура
type stack struct {
	items []rune
}

// push Записываем значение в стэк
func (s *stack) push(i rune) {
	s.items = append(s.items, i)
}

// pop достаем значение из стэка
func (s *stack) pop() rune {
	l := len(s.items) - 1
	lastItem := s.items[l]
	s.items = s.items[:l]
	return lastItem
}
