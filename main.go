package main

import (
	"fmt"
	"sync"
)

// square печатает в терминал квадрат числа
func square(n int, wg *sync.WaitGroup) {
	fmt.Printf("%d^2 = %d\n", n, n*n)
	wg.Done() // уменьшаем счетчик ожидаемых горутин
}

//
//func main() {
//	var wg sync.WaitGroup
//	go func() {
//		wg.Add(1)
//		defer wg.Done()
//		fmt.Println("1st goroutine sleeping...")
//		time.Sleep(1)
//	}()
//	wg.Wait()
//	fmt.Println("All goroutines complete.")
//}

//func main()  {
//	var wg sync.WaitGroup
//
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//
//		fmt.Println("1st goroutine sleeping...")
//		time.Sleep(100 * time.Millisecond)
//	}()
//
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//
//		fmt.Println("2nd goroutine sleeping...")
//		time.Sleep(200 * time.Millisecond)
//	}()
//
//	wg.Wait()
//	fmt.Println("All goroutines complete.")
//}

func main() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}

	const numGreeters = 5
	var wg sync.WaitGroup

	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}

	wg.Wait()
}
