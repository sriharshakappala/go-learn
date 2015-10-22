package main

import (
  "fmt"
)

func first() {
  fmt.Println("1st")
}

func second() {
  fmt.Println("2nd")
}

func main() {
  defer second()
  first()
}

// Defer usually takes the funcion call to the end of the block
/*
Example:
f, _ := os.Open(filename)
defer f.Close()
Refer screenshot for other advantages
*/
