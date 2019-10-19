package main

import (
    "fmt"
)

func main() {
    x := 378
    y := 273
    d := TrivialGCD(x, y)
    fmt.Println(d)
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

func Min2a(a,b int) int {
    if a > b {
      return b
    }
    return a
}
