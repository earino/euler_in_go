package main

import (
        "fmt"
)

func generate(n uint64, p bool) int {
  var counter int = 0
  if p {
    fmt.Printf(" %d ->", n)
  }

  for ;n != 1 && counter < 100000; {
    counter++
    if n % 2 == 0 {
      n = n / 2
    } else {
      n = 3 * n + 1
    }
    if p {
      if n == 1 {
        fmt.Printf(" %d", n)
      } else {
        fmt.Printf(" %d ->", n)
      }
    }
  }

  if p {
    fmt.Println("")
  }

  return counter
}

func main() {
  var i uint64= 1
  var largest, tmp = 0, 0
  var largest_at uint64 = 0

  for ;i < 1000000 ;i++ {
    tmp = generate(i, false)
    if tmp > largest {
      largest = tmp
      largest_at = i
    }
  }

  generate(largest_at, true)
}
