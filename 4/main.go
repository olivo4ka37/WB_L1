package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Пул воркеров (количество воркеров)
	var workerPool int

	// declareWorker Пользователь задает значение пулу воркеров через терминал
	declareWorker(&workerPool)

	// chanStr Канал для записи данных из главного потока
	chanStr := make(chan string)

	// Бесконечный цикл при работе которого мы будем иметь главный поток
	// и воркеров которые будут читать данные из него.
	for workerPool > 0 {
		// Горутина которая генерирует строку и заносит результат в канал
		go generateString(chanStr)
		// Цикл вызова воркеров кол-ву равному workerPool
		for i := 0; i < workerPool; i++ {
			go worker(i+1, chanStr)
		}
		time.Sleep(time.Second)
	}

}

// declareWorker Считываем с консольного ввода значение воркеров
func declareWorker(w *int) {
	in := bufio.NewReader(os.Stdin)

	fmt.Println("Write amount of workers")

	fmt.Fscan(in, w)
}

// worker Считывает значение из канала
func worker(id int, str chan string) {
	time.Sleep(time.Millisecond * 100)
	fmt.Printf("Worker %d, readed: %s\n", id, <-str)
}

// generateString Генерирует случайную строку в 8 символов и записывает её в канал
func generateString(str chan string) {
	time.Sleep(time.Millisecond * 100)
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	specials := "~=+%^*/()[]{}/!@#$?|"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		digits + specials
	length := 8
	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	result := string(buf)

	str <- result
}
