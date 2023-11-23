package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Структура для использования WaitGroup в функциях
type w8Group struct {
	wg *sync.WaitGroup
}

// Горутина проверяет закрылся ли канал пытаясь вычитать из него данные,
// если вернулось стандартное значение (не ок), то завершаем функцию
func (wg *w8Group) okChanel(ch <-chan int) {
	defer wg.wg.Done()

	for {
		x, err := <-ch
		if !err {
			fmt.Println("goroutine checking if channel closed finished")
			return
		}
		fmt.Printf(" The number is: %d\n", x)
	}
}

// range завершится сам, как только канал закроется
func (wg *w8Group) rangeChanel(ch <-chan int) {
	defer wg.wg.Done()

	for x := range ch {
		fmt.Printf(" The number is: %d\n", x)
	}
}

// используем отдельный канал для завершения горутины
func (wg *w8Group) stopChanel(stop <-chan struct{}, ch <-chan int) {
	defer wg.wg.Done()

	for {
		select {
		case v := <-ch:
			fmt.Printf(" got: %d\n", v)
		case <-stop:
			fmt.Println("goroutine with cancel channel finished")
			return
		}
	}
}

// используем сигнал из контекста для завершения горутины
func (wg *w8Group) doneChanel(ctx context.Context, ch <-chan int) {
	defer wg.wg.Done()

	for {
		select {
		case v := <-ch:
			fmt.Printf(" got: %d\n", v)
		case <-ctx.Done():
			fmt.Println("close context goroutine finished")
			return
		}
	}
}

// используем сигнал из контекста для завершения горутины
// в данном случае контекст завершается по таймауту, так как выполняется долгая операция
func (wg *w8Group) timeoutChanel(ctx context.Context, ch <-chan int) {
	defer wg.wg.Done()

	for {
		select {
		case <-time.After(10 * time.Second): // long operation
			fmt.Printf(" got: %d\n", <-ch)
		case <-ctx.Done():
			fmt.Println("context timeout goroutine finished")
			return
		}
	}
}

// запись случайных чисел в канал с определенной периодичностью
// завершается по сигналу контекста

func (wg *w8Group) tickerChanel(ctx context.Context, ch chan int) {
	defer wg.wg.Done()

	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker.C:
			a := rand.Int()
			ch <- a
			fmt.Printf("sent: %d\n", a)
		case <-ctx.Done():
			fmt.Println("exiting from writer")
			ticker.Stop()
			return
		}
	}
}

func main() {
	// context timeout
	fmt.Println("Context timeout exceeded")
	ctx, cancel := context.WithCancel(context.Background())
	timeout, _ := context.WithTimeout(context.Background(), 5*time.Second)
	ch := make(chan int, 1)
	g := w8Group{
		wg: &sync.WaitGroup{},
	}

	g.wg.Add(1)
	go g.tickerChanel(ctx, ch)
	go g.timeoutChanel(timeout, ch)

	g.wg.Wait()
	cancel()

	// canceling context
	fmt.Println("\nContext cancel")
	ctx, cancel = context.WithCancel(context.Background())
	ch = make(chan int, 1)

	g.wg = &sync.WaitGroup{}
	g.wg.Add(2)
	go g.tickerChanel(ctx, ch)
	go g.doneChanel(ctx, ch)

	time.Sleep(3 * time.Second)
	cancel()
	close(ch)
	g.wg.Wait()

	// stopping channel
	fmt.Println("\nStop channel")
	ch = make(chan int, 1)
	stop := make(chan struct{})
	ctx, cancel = context.WithCancel(context.Background())

	g.wg = &sync.WaitGroup{}
	g.wg.Add(2)
	go g.tickerChanel(ctx, ch)
	go g.stopChanel(stop, ch)

	time.Sleep(3 * time.Second)
	stop <- struct{}{}
	cancel()
	close(ch)
	g.wg.Wait()

	// ranging channel
	fmt.Println("\nRanging channel")
	ch = make(chan int, 1)
	ctx, cancel = context.WithCancel(context.Background())

	g.wg = &sync.WaitGroup{}
	g.wg.Add(2)
	go g.tickerChanel(ctx, ch)
	go g.rangeChanel(ch)

	time.Sleep(3 * time.Second)
	close(ch)
	cancel()
	g.wg.Wait()

	// closing channel
	fmt.Println("\nClosing channel")
	ch = make(chan int, 1)
	ctx, cancel = context.WithCancel(context.Background())

	g.wg = &sync.WaitGroup{}
	g.wg.Add(2)
	go g.tickerChanel(ctx, ch)
	go g.okChanel(ch)

	time.Sleep(3 * time.Second)
	close(ch)
	cancel()
	g.wg.Wait()
}
