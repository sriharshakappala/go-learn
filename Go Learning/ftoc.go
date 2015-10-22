package main

import (
  "fmt"
)

func main()  {
  fmt.Println("Enter the temperature in fareinheit:");
  var fareinheit float64;
  fmt.Scanf("%f", &fareinheit);
  var celsius float64
  // var conversion_rate float64 = 5/9;
  // fmt.Println(conversion_rate);
  celsius = (fareinheit - 32) * 0.5555;
  fmt.Println("The temperature in celsius is:", celsius);
}
