package main

import "fmt"

func main() {
  const x string = "Hello World"
  fmt.Println(x)
  // Doing below will result error
  // x = "Hello New World"
  // fmt.Println(x)
}
