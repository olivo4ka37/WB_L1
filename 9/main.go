package main

import "fmt"

func main() {
	xArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	readChan := make(chan int)
	writeChan := make(chan int)

	go writer(xArray, writeChan)
	go reader(writeChan, readChan)

	// Выводит все значения из канала
	for x := range writeChan {
		fmt.Println(x)
	}

}

// writer Записывает в канал все значения из массива
func writer(arr []int, writeChan chan int) {
	for _, x := range arr {
		writeChan <- x
	}

	close(writeChan)
}

// reader Считывает значения из writeChan возводит во вторую степень, потом записывает в readChan
func reader(writeChan chan int, readChan chan int) {
	for {
		x, ok := <-writeChan

		// завершается в случае закрытия 1 канала
		if ok != true {
			close(writeChan)
			return
		}

		readChan <- x * x

	}
}
