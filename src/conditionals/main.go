package main

import (
  "fmt"
)

func main() {

}

func SimpleFunction(x, y int) int {
    if x == y {
      return 0
    } else if x > y {
      return 1
    } else { // we can infer x < y
      return -1
    }
}

func SameSign(x, y int) bool {
    if x * y >= 0 { // >= is "greater than or equal"
      return true
    }
    // if we make it here, we know that x*y < 0
    return false
}

// when we declare variable, it has a 'scope' (lifetime) that exists inside
// of the 'block' in which it lives

func PositiveDifference(a, b int) int {
    var c int
    if a == b {
      c = 0
    } else if a > b {
        c = a-b
    } else { // a < b
        c = b-a
    }
    return c
}

// Other conditions
/*
>, <, >=, <=, ==

"not equal to" !=
"not" !x (this evaluates to true when x is false and false if x is true)
*/
