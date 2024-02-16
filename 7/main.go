package main

import (
	"fmt"
	"sync"
)

// Структура для использования WaitGroup и Mutex в функциях
type w8Group struct {
	wg sync.WaitGroup
	m  sync.Mutex
}

func main() {
	w := w8Group{}
	c := map[int]int{}

	// Вызов Горутин равное количеству n
	n := 200
	w.wg.Add(n)
	for i := 0; i < n; i++ {
		go w.concurentWriter(c, i)
	}
	w.wg.Wait()

	// Вывод длины map (Проверка результата)
	fmt.Println(len(c))
}

// Конкурентная запис в map структуру с использованием Mutex
func (w *w8Group) concurentWriter(m map[int]int, i int) {
	w.m.Lock()
	m[i] = i
	w.m.Unlock()
	w.wg.Done()

}
