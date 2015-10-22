package main

import (
  "fmt"
)

func main() {
  x, err := fmt.Println("Hello!");
  if err == nil {
    fmt.Println(x);
    fmt.Println(err);
  } else {
    fmt.Println(err);
  }
  // else part must start with the same line if ended
}
