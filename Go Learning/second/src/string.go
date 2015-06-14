package main

import "fmt"

func main() {
  fmt.Println(len("Hello World"))
  fmt.Println("Hello World"[1])
  fmt.Println("Hello " + "World")
}

// Strings are indexed from 0
// Character is represented as a byte. So "Hello World"[1] gives 101 instead of 'e' - ASCII code
