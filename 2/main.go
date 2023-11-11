package main

import (
	"fmt"
	"time"
)

func main() {
	arr1 := []int{2, 4, 6, 8, 10}

	for i := 0; i < len(arr1); i++ {
		go func(i int) {
			fmt.Println(arr1[i] * arr1[i])
		}(i)
	}

	time.Sleep(100 * time.Millisecond)
}
