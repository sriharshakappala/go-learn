package main

import "fmt"

func main() {
  var x string = "hello"
  var y string = "hello"
  var z string = "world"
  fmt.Println(x == y)
  fmt.Println(x == z)
  another()
}

func another() {
  var(
    x = "hello"
    y = "hello"
    z = "world"
    )
  fmt.Println(x == y)
  fmt.Println(x == z)
}

// another() function shows the alternate way of declaring variables
