package main

import(
  "fmt"
)

func main() {
  fmt.Println("Enter a number:")
  var number float64;
  fmt.Scanf("%f", &number);
  output := number * 2
  fmt.Println("Number after doubling is:", output)
}
