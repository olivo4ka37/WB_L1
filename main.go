package main

import "fmt"

func main() {
	cat := Cat{}
	dog := Dog{}
	cow := Cow{}

	fmt.Println(cat.Voice()) // Мяу
	fmt.Println(dog.Voice()) // Гав
	fmt.Println(cow.Voice()) // Мууу
}

type IVoiceable interface {
	Voice() string
}

type Cat struct {
	// …
}

func (c Cat) Voice() string {
	return "Meow"
}

type Cow struct {
	// …
}

func (c Cow) Voice() string {
	return "Mooooo"
}

type Dog struct {
	// …
}

func (c Dog) Voice() string {
	return "Gap"
}

// BEGIN (write your solution here)

// END
