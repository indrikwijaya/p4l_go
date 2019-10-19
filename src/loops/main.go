package main

import (
  "fmt"
)

func main() {
  fmt.Println("Loops.")

  n := 5
  m := Factorial(n)

  fmt.Println(m)

  fmt.Println(Factorial(-100))
}

// first, a factorial functions

func Factorial(n int) int{
    // let's handle negative inputs
    if n < 0 {
      // panic() will print an error message and immediately end program.
      panic("Error: Negative input given to Factorial()")
    }
    p := 1
    i := 1
    // Go doesn't have a while keyword. They use "for"
    for i <= n {
        p = p*i
        i = i+1
    }

    return p
}

// Exercise: Write a function in Go using a while loop that takes an integer n
// and returns the sum of the first n positive integers.

func SumFirstNIntegers(n int){
  // let's handle negative inputs
  if n < 0{
    panic("Error: Negative input given")
  }
  sum := 0
  var i int = 1

  for i <= n {
    sum += i // sum = sum + i
    i ++ // i = i++, there is also i-- which means i = i-1
  }
  
  return sum
}
