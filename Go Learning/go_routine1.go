package main

import (
  "fmt"
)

func f(n int) {
  for i := 0; i < 1000; i++ {
    fmt.Println(n, ":", i)
  }
}

// func main() {
//   go f(0)
//   var input string
//   fmt.Scanln(&input)
// }

func main() {
  for i := 0; i < 10; i++ {
    go f(i)
  }
  var input string
  fmt.Scanln(&input)
}
