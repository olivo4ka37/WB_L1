package main

import "fmt"

func main() {

	fmt.Println(reverseWords("snow    dog sun"))
}

// reverseWords пишет слова из строки в обратном порядке, игнорирует мультипробелы
func reverseWords(s string) string {
	result, word := "", ""

	for i := 0; i < len(s); i++ {
		// i== len(s)-1 Проверяет конецу строки, для того чтобы добавить последнее слово.
		if i == len(s)-1 {
			// Проверяет необходимость для того чтобы добавить последний символ в переменную word.
			if string(s[i]) != " " {
				word += string(s[i])
			}
			result = word + result

			// string(s[i]) == " " && string(s[i+1]) != " " && word != "" Проверяет на пробелы между словами,
			// игнорирует мультипробелы.
			// После всех проверок добовляет символ если он у нас есть.
		} else if string(s[i]) == " " && string(s[i+1]) != " " && word != "" {
			result = " " + word + result
			word = ""
		} else if string(s[i]) != " " {
			word += string(s[i])
		}
	}

	return result

}
