package main

import "fmt"

func main() {
  var tmp, sum, prev, cur int = 0, 0, 1, 2

  for   ; cur < 4000000; {
    fmt.Printf("fibonacci term: %d\n", cur)
    if cur % 2 == 0 {
      sum += cur 
    }
    tmp = cur
    cur = cur + prev
    prev = tmp
    
  }
  fmt.Printf("sum is: %d\n", sum)
}
