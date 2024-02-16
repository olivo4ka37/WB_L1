package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Структура для использования WaitGroup в функциях
type w8Group struct {
	wg *sync.WaitGroup
}

// stopChannel Данная функция останавливается как только в канал приходит любое значение.
// (Какое значение будет переданно в канал и какой тип данных будет неважно, как пример я использовал канал типа bool)
func (w *w8Group) stopChanel(stop chan bool) {
	w.wg.Done()

	for {
		select {
		case <-stop:
			fmt.Println("Received value from chanel. Exiting from stopChanel...")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}

// stopContext Данная функци останавливается как только канал из контекста закрывается.
func (w *w8Group) stopContext(ctx context.Context) {
	w.wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Received stop signal. Exiting from stopContext...")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}

// stopEndChanel Данная функция завершает свою работу, когда канал закрывается.
func (w *w8Group) stopEndChanel(ch <-chan int) {
	defer w.wg.Done()
	for {
		_, ok := <-ch
		if ok != true {
			fmt.Println("Chanel is ended. Exiting from stopEndChanel...")
			return
		}
		fmt.Println("Working...")
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	w := w8Group{
		wg: &sync.WaitGroup{},
	}
	w.wg.Add(3)

	// Блок кода в котором запускаем горутину stopChanel, а потом останавливаем её.
	stop := make(chan bool)
	go w.stopChanel(stop)

	time.Sleep(1 * time.Second)
	fmt.Println("Sending value to chanel...")
	stop <- true

	// Блок кода в котором запускаем горутину stopChanel, а потом останавливаем её.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go w.stopContext(ctx)

	time.Sleep(1 * time.Second)
	fmt.Println("Cancelling the worker...")
	cancel()

	// Блок кода в котором запускаем горутину stopEndChanel

	ch := make(chan int)

	go w.stopEndChanel(ch)

	for i := 0; i < 10; i++ {
		ch <- 1
	}
	close(ch)

	// Демонстрация завершения всех горутин
	time.Sleep(2 * time.Second)
	fmt.Println("Main function exits.")

	w.wg.Wait()

}
