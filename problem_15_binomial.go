package main

import ("fmt"
        "math/big")

const PATHS int64 = 137846528820
const MAX = 20

var path_count = 0

func factorial(n *big.Int) (result *big.Int) {
  result = new(big.Int)

  switch n.Cmp(&big.Int{}) {
    case -1, 0:
      result.SetInt64(1)
    default:
      result.Set(n)
      var one big.Int
      one.SetInt64(1)
      result.Mul(result, factorial(n.Sub(n, &one)))
  }
  return
}
    
func main() {
  var numerator = factorial(big.NewInt(2 * MAX))
  var denominator = factorial(big.NewInt(MAX)).Mul(factorial(big.NewInt(MAX)), factorial(big.NewInt(MAX)))

  fmt.Printf("%d / %d = %d\n", numerator, denominator, numerator.Div(numerator, denominator))
}
