package main

import (
  "fmt"
)
func main()  {
  var number byte;
  fmt.Scanf("%i", &number)
  switch number {
  case 0: fmt.Println("Zero")
  case 1: fmt.Println("One")
  case 2: fmt.Println("Two")
  default: fmt.Println("Some other fucking number")
  }
}
