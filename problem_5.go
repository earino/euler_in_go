package main

import "fmt"

func is_it_divisible_by(is_it int, from int, to int) bool {
  var i int = from;
  for ; i < to; i ++ {
    if is_it % i != 0 {
      return false
    }
  } 
  return true
}

func main() {
  var x int = 1

  for ; ; x++ {
    if is_it_divisible_by(x, 1, 20) {
      fmt.Printf("%d\n", x)
      return
    }
  }
}
