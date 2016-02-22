package main

import "fmt"

type Salutation struct {
  name string
  greeting string
}

const (
  PI = 3.14
  Language = "Go"
)

const (
  A = iota
  B
  C
)

func Greet(salutation Salutation) {
  _, alternateMessage := CreateMessage(salutation.name, salutation.greeting)
  fmt.Println(alternateMessage)
}

func CreateMessage(name, greeting string) (message string, alternate string) {
  message = greeting + " " + name
  alternate = "HEY! "  + name
  return
}

func main()  {
  var s = Salutation { "John", "Hello!" }
  Greet(s)

  fmt.Println(PI)
  fmt.Println(Language)
  fmt.Println(A, B, C)
}
