package main

import (
	"fmt"
	"sync"
	"sync/atomic"
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
			s.atomicIncrease()
		}(i)
	}
	s.wg.Wait()

	fmt.Println(s.count)

}

type schetchik struct {
	count int64 //Значение в счетчике
	wg    sync.WaitGroup
	mut   sync.Mutex
}

// worker инкрементирует значение счетчика в структуре на 1
// Использует методы синхронизации Mutex и WaitGroup
func (s *schetchik) worker() {
	s.mut.Lock()
	s.count++
	s.mut.Unlock()
	s.wg.Done()
}

// atomicIncrease Инкрементирует значение счетчика в структуре на 1
// Использует методы синхронизации Atomic и WaitGroup
func (s *schetchik) atomicIncrease() {
	atomic.AddInt64(&s.count, 1)
	s.wg.Done()
}

// Я решил реализовать два способа так как, Mutex под капотом содержит atomic.
