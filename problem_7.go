package main

import (
        "fmt" 
)

func is_prime(x int64) bool {
  var i int64 = 2

  for ; i < x / 2; i++ {
    if x % i == 0 {
      return false
    }
  }
  return true
}

func main() {
  var i, prime_counter int = 2, 0

  for ; ; i++ {
    if is_prime(int64(i)) {

      if prime_counter == 10001 {
        fmt.Printf("%d\n", i)
        return
      }
      prime_counter++
    } 
  }
}
