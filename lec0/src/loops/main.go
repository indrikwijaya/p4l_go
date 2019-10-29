package main

import (
  "fmt"
)

func main() {
  fmt.Println("Loops.")

  n := 5
  m := Factorial(n)

  fmt.Println(m)

  // fmt.Println(Factorial(-100))

  // fmt.Println(SumFirstNIntegers(10))

  // closing: infinite loop in disguise
  var i uint = 10

  for ; i >= 0; i -- {
    fmt.Println(i)
  }
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

    // i still lives here and I could use it
    return p
}

func AnotherFactorial(n int) int {
    if n < 0 {
      // panic() will print an error message and immediately end program.
      panic("Error: Negative input given")
    }
    p := 1

    // for every integer i between 1 and n, p = p*i
    for i := 1; i <= n; i++ { // i has scope up to end of for loop
      // or i := n; i >= 1; i-- but this doesn't really make sense/encouraged
      p *= i
    }

    // i does not live here, which is good

    return p
}

// Exercise: Write a function in Go using a while loop that takes an integer n
// and returns the sum of the first n positive integers.
/*
func SumFirstNIntegers(n int){
  sum := 0
  var i int = 1
  for i <= n {
    sum += i // sum = sum + i
    i++ // i = i++, there is also i-- which means i = i-1
  }

  return sum
}
*/

func AnotherSum(n int) int {
    sum := 0

    for k := 1; k <= n; k++ {
      sum += k
    }

    return sum
}

// Exercise: write a function SumEven that sums all even numbers up to and
// possibly including n.

/*
func SumEven(k int){
    sum := 0

    for j := 2; j <= k; j = j + 2 {
      sum += j
    }

    return sum
}
*/
