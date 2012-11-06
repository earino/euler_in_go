package main

import (
        "fmt" 
        "math/big"
)

func generate_sieve(max int) [2000000]bool {
  var i int64 = 3;
  var retval [2000000]bool

  retval[0] = false;
  retval[1] = false;
  retval[2] = true;

  for ; i < int64(max); i++ {
    retval[i] = true
  }
  i = 2
  for ; i < int64(max); i++ {
    if retval[i] == true {
      var j int64 = i * i
      for ; j < int64(max); {
        retval[j] = false
        j = j + i
      }
    }
  }

  return retval
}

func main() {
  var i, max int = 0, 2000000
  var primes [2000000]bool = generate_sieve(max)
  var sum = new(big.Int)
  
  for ; i < max; i++ {
    if primes[i] == true {
      sum = sum.Add(sum, big.NewInt(int64(i)))
    }
  }

  fmt.Printf("%d\n", sum)
}
