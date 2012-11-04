package main

import "fmt"

func is_palindrome(data string) bool {
  var i, length int = 0, len(data)

  for ; i < length / 2; i++ {
    if data[i] != data[len(data) - 1 - i] {
      return false
    }
  }

  return true
}

func main() {
  var x int64 = 100
  var largest int64 = 0

  for ; x < 1000; x++ {
    var y int64 = 100
    for ; y < 1000; y++ {
      var data string = fmt.Sprintf("%d", x * y)
      if is_palindrome(data) {
        if x * y > largest {
          largest = x * y
        }
      }
    }
  }
  fmt.Printf("largest palindrone: %d\n", largest)
}
