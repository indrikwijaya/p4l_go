package main

import (
    // "fmt"
    "time"
    "log"
)

func main() {
    x := 378202873
    y := 273147834

    start := time.Now()
    TrivialGCD(x, y)
    elapsed := time.Since(start)
    log.Printf("TrivialGCD took  %s", elapsed)

    start2 := time.Now()
    EuclidGCD(x, y)
    elapsed2 := time.Since(start2)
    log.Printf("EuclidGCD took %s", elapsed2)


}

/*
TrivialGCD(a,b)
    d <- 1
    m <- Min2(a,b)
    for every integer p from 1 to m
      if p is a divisor of a and b
        d <- p
    return d
*/

func TrivialGCD(a, b int) int {
    d := 1
    m := Min2a(a, b)
    for p := 1; p <= m; p ++ {
      // remainder operation Remainder(n, k) is n%k eg (14%3 = 1)
      if a % p == 0 && b % p == 0 { // if statement is only true if both are true
            // if i'm here, i know that the answer to both is "YES"
              d = p
        }
    }

    return d
}

// there is an "OR" of 2 statements as well, operator is ||
// e || f is true if one or both of e and f is true. It's only false if
// both e and f are false

func AnotherSumEven(n int) int {
    sum := 0

    for i := 2; i <= n; i++ {
      // is i even?
      if i%2 == 0 {
          sum += i
        }
    }

    return sum
}

/*
EuclidGCD(a, b)
  while a not equal to b
    if a > b
        a = a - b
    else // b > a
        b = b - a
  return a
*/

func EuclidGCD(a, b int) int {
    for a != b {
        if a > b {
            a = a - b
        } else {
            b = b - a
        }
    }
    return a
}

func Min2a(a,b int) int {
    if a > b {
      return b
    }
    return a
}
