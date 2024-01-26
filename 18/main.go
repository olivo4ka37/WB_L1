package main

import (
	"fmt"
	"sync"
)

func main() {
	s := schetchik{} // Декларация структуры счетчика

	amountOfWorkers := 10000 // Декларация количества воркеров

	// Цикл вызова воркеров кол-ву равному amountOfWorkers
	for i := 0; i < amountOfWorkers; i++ {
		// Увеличиваем счетчик WaitGroup на 1
		s.wg.Add(1)
		//Вызываем горутину инкриментирующую счетчик в нашей структуре на 1
		go func(i int) {
			s.worker()
		}(i)
	}
	s.wg.Wait()

	fmt.Println(s.count)
}

type schetchik struct {
	count int //Значение в счетчике
	wg    sync.WaitGroup
	mut   sync.Mutex
}

// worker инкриментирует значение счетчика в структуре на 1
// Использует методы синхронизации Mutex и WaitGroup
func (s *schetchik) worker() {
	s.mut.Lock()
	s.count++
	s.mut.Unlock()
	s.wg.Done()
}
