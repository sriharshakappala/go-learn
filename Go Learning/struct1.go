package main

import (
  "fmt"
)

type Rectangle struct {
  x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
  l := r.x2 - r.x1
  w := r.y2 - r.y1
  return l * w
}

func main() {
  r := Rectangle{ 0, 0, 10, 10 }
  fmt.Println(r.area())
}
