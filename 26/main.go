package main

import (
	"fmt"
	"strings"
)

func main() {
	// Тест1
	fmt.Println(UniqueString("abcd"))
	// Тест2
	fmt.Println(UniqueString("abCdefAaf"))
	// Тест3
	fmt.Println(UniqueString("aabcd"))

}

var exists = struct{}{}

// Set Структура множества значений типа string
type Set struct {
	m map[string]struct{}
}

// NewSet Создает новое множество значений типа string
func (Set) NewSet() *Set {
	s := &Set{}
	s.m = make(map[string]struct{})
	return s
}

// add Добавляет элемент в множество
func (s *Set) add(value string) {
	s.m[value] = exists
}

// contains Проверяет есть ли этот элемент в множестве
func (s *Set) contains(value string) bool {
	_, c := s.m[value]
	return c
}

// UniqueString Проверяет строку на уникальность символовов регистр букв не учитывается.
func UniqueString(str string) bool {
	var s Set
	set := s.NewSet()
	for i := 0; i < len(str); i++ {
		xLetter := strings.ToLower(string(str[i]))
		if set.contains(xLetter) {
			return false
		}

		set.add(xLetter)
	}

	return true
}
