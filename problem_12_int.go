package main

import (
        "log"
)

var numbers map[uint64]uint64

func triangle(up_to uint64) uint64 {
  numbers[up_to] = 0

  var term = up_to
  var other_term = numbers[up_to - 1]

  numbers[up_to] = term + other_term
  return numbers[up_to]
}

/* first algorithm, so very slow */
func count_divisors(value uint64) int {
  var i uint64 = 1
  var retval int = 2
  
  for ; i < value/2; i++ {
    var is_evenly_divisible = value % i == 0

    if is_evenly_divisible {
      retval++ 
    }
  }

  return retval
}

/* much better algorithm, prime factorization is a hell
    of a drug */
func divisors(x uint64) int {
    var limit uint64 = x 
    var numberOfDivisors int = 0 
    var i int = 1

    for ; uint64(i) < limit; i++  {
        if x % uint64(i) == 0  {
            limit = x / uint64(i) 
            if  limit != uint64(i)  {
                numberOfDivisors++ 
            }
            numberOfDivisors++ 
        }
    }

    return numberOfDivisors 
}


func main() {
  var i, largest int = 2, 0

  numbers = make(map[uint64]uint64)

  numbers[1] = uint64(1)
  numbers[2] = uint64(3)
  numbers[3] = uint64(6)

  for ; ; i++ {
    var value = triangle(uint64(i))
    var divisors = divisors(value)

    if  divisors > largest  {
      largest = divisors
      log.Printf("value: %d %d\n", value, largest)
    }
  }
}
