package main

import "fmt"

func main() {
  var sum, i int = 0, 1

  for   ; i < 1000; i++  {
    if i % 5 == 0 || i % 3 == 0 {
      sum += i
    }
  }
  fmt.Printf("sum is: %d\n", sum)
}
