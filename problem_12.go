package main

import (
        "math/big"
        "log"
)

var numbers map[uint64]*big.Int

func triangle(up_to uint64) *big.Int {
  numbers[up_to] = big.NewInt(0)

  var term = big.NewInt(int64(up_to))
  var other_term = numbers[up_to - 1]

  //log.Printf("%d %d %d\n", up_to, term, other_term)

  var retval = numbers[up_to].Add(term, other_term)
  numbers[up_to] = retval
    
  return retval
}

func count_divisors(value *big.Int) int {
  var i = big.NewInt(1)
  var retval int = 1
  
  for ; i.Cmp(value) == -1; i.Add(i, big.NewInt(1)) {
    var mod = new(big.Int)
    mod.Mod(value, i)

    var zero = big.NewInt(0)
    var is_evenly_divisible = zero.Cmp(mod) == 0

    if (is_evenly_divisible) {
      retval++ 
    }
  }

  return retval
}

func main() {
  var i, largest int = 2, 0

  numbers = make(map[uint64]*big.Int)
  numbers[1] = big.NewInt(1)
  numbers[2] = big.NewInt(3)
  numbers[3] = big.NewInt(6)

  for ; ; i++ {
    var value = triangle(uint64(i))
    var divisors = count_divisors(value)

    if  divisors > largest  {
      largest = divisors
      log.Printf("value: %d %d\n", value, largest)
    }
  }
}
