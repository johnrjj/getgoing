package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

type Printer func(string) // doesnt return anything

const (
	PI       = 3.14
	Language = "Go"
)

const (
	A = iota
	B
	C
)

func Greet(salutation Salutation, do Printer) {
	message, alternateMessage := CreateMessage(salutation.name, salutation.greeting, "yo")
	do(message)
	do(alternateMessage)
}

func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	fmt.Println(len(greeting))
	message = greeting[0] + " " + name
	alternate = "HEY! " + name
	return
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

func main() {
	var s = Salutation{"John", "Hello!"}
	Greet(s, CreatePrintFunction("!!!"))

	fmt.Println(PI)
	fmt.Println(Language)
	fmt.Println(A, B, C)
}
