package main

import "fmt"

func main() {
  var largest, d, n int64 = 0, 2, 600851475143

  for ; n > 1; {
    for ; n % d == 0; {
      if d > largest {
        largest = d
      }
      n /= d
    }
    d += 1
  }
  fmt.Printf("%d\n", largest)
}
