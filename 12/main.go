package main

import "fmt"

func main() {
	// Массив строк
	xArray := []string{"cat", "cat", "dog", "cat", "tree"}
	// Множество значений string
	stringSet := make(map[string]struct{})

	// Проверяем наличие значений из xArray в stringSet, если нету записываем в stringSet
	for _, x := range xArray {
		_, contains := stringSet[x] //contains - булевое значение определяющее наличие данного ключа в множестве
		if !contains {
			stringSet[x] = struct{}{}
		}
	}

	fmt.Println(stringSet)

}
