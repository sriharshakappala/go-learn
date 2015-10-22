package main

import (
  "fmt"
)

func main() {
  var x [5]float64;
  x[0] = 98;
  x[1] = 89;
  x[2] = 89;
  x[3] = 89;
  x[4] = 89;

  var total float64 = 0;

  for i := 0; len(x); i++ {
    total += x[i]
  }

  fmt.Println(total/float64(len(x)));
}
