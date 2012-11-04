package main

import "fmt"

func sum_of_squares(x int64) int64 {
  var i, sum int64 = 0, 0
  for ; i <= x; i++ {
    sum += i * i
  }

  return sum
}

func square_of_sums(x int64) int64 {
  var i, sum int64 = 0, 0
  for ; i <= x; i++ {
    sum += i
  }

  return sum * sum
}

func main() {
  fmt.Printf("difference: %d\n", square_of_sums(100) - sum_of_squares(100))
}
