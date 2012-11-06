package main

import (
        "fmt" 
        "math"
)


func compute_triplet(a int, b int) int {
  var tmp float64 = math.Sqrt(float64(a * a + b * b))

  if math.Floor(tmp) == math.Ceil(tmp) {
    return int(tmp)
  }
  //var c int = int(math.Sqrt(float64(a * a + b * b)))

  return -1
}


func main() {
  var a, b, c int = 0, 0, 0

  for a = 1; a < 1000; a++ {
    for b = a+1; b < 1000; b++ {
      c = compute_triplet(a, b)
      
      if c != -1 && a + b + c == 1000 {
        fmt.Printf("%d^2 + %d^2 = %d^2\n", a, b, c)
      }
    }
  }
}
