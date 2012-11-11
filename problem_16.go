package main

import ("fmt"
        "strconv"
        "math/big")

func main() {
  var result = big.NewInt(2).Exp(big.NewInt(2), big.NewInt(1000), nil)
  var result_as_string = result.String()
  var sum int = 0
  var i int = 0

  for ; i < len(result_as_string); i++ {
    addend, err := strconv.Atoi(string(result_as_string[i]))
    if err != nil {
      fmt.Printf("err: %s", err)
    }
    sum += addend
  }

  fmt.Printf("result: %d\n", result)
  fmt.Printf("sum: %d\n", sum)
}
