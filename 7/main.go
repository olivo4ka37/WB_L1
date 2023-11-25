package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	myMap := make(map[string]bool)
	wg := sync.WaitGroup{}

	ch := make(chan string)
	wg.Add(6)
	go writer(myMap, ch, &wg)
	go printer(myMap, ch, &wg)

	wg.Wait()
}

func writer(mp map[string]bool, ch chan string, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	for {
		newStr := generateString()
		ch <- newStr
		mp[newStr] = true

		time.Sleep(time.Second)

		wg.Done()
	}
}

func printer(mp map[string]bool, ch chan string, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)
	for {
		fmt.Println(mp[<-ch])

		time.Sleep(time.Second)

		wg.Done()
	}
}

// generateString Генерирует случайную строку в 8 символов и записывает её в канал
func generateString() string {
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
	time.Sleep(time.Second)

	return result
}
