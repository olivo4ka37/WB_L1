package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// chanStr Канал для записи данных из главного потока
	chanStr := make(chan string)

	var amountOfTime int

	declareTime(&amountOfTime)

	// Горутина которая генерирует строку и заносит результат в канал
	go generateString(chanStr)
	// Горутина которая читает строку из канала и выводи в std	out
	go reader(chanStr)

	// Время работы программы
	time.Sleep(time.Second * time.Duration(amountOfTime))
}

// generateString Генерирует случайную строку в 8 символов и записывает её в канал
func generateString(str chan string) {
	for {
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
		time.Sleep(time.Millisecond * 500) // Замедляем для удобного отслеживания работы
	}
}

// reader Считывает значение из канала
func reader(str chan string) {
	for {
		fmt.Printf("%s\n", <-str)
		time.Sleep(time.Millisecond * 500) // Замедляем для удобного отслеживания работы
	}
}

// declareTime Считываем с консольного ввода значение количества секунд
func declareTime(w *int) {
	in := bufio.NewReader(os.Stdin)

	fmt.Println("Write how much seconds program should work\n")

	fmt.Fscan(in, w)
}
