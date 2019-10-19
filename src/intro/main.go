package main

import(
  "fmt"
)
// Go won't read this line

/*
Everthing here won't be read (multiline comments)

Today: 4 variable types
int: integers
uint: unsigned integers (non-negative)
bool: True/False Boolean variable
float64: decimal number

next time:
byte: single symbol
string: contiguous collection of symbols (words)
*/

func main() {
  fmt.Println("Let's get started!")

  var j int = 14 // j is an integer variable
  var x float64 = -2.3
  var yo string = "Hi" // default is empty string ""
  var u uint = 14
  var symbol byte = 'H' // prints 72?
  var statement bool = true // defaults to false

  fmt.Println(j)
  fmt.Println(x)
  fmt.Println(yo)
  fmt.Println(u)
  fmt.Println(symbol)
  fmt.Println(statement)

  // shorthand declarations

  i := -6 // automatically type int (if need uint, need to declare)
  hi := "Yo " // Go automatically gives this type string
  k := 34 // automatically an int
  // secondstatement := true

  // arithmetic on numeric variables

  fmt.Println((i+j)*2*k)
  fmt.Println(2*x - 3.16)

  fmt.Println(hi + yo)

  fmt.Println(k/j) // Go views as integer division ... throws away remainder

  // if we want actual value of k/j, we use type conversions
  fmt.Println(float64(k)/float64(j))

  // not all type conversions will work

  /*
  var ok bool = bool(0) // false?
  fmt.Println(ok)
  */

  var p int = -187
  var s uint = uint(p)

  fmt.Println(s)

  m := 9223372036854775807
  fmt.Println(m+1)

  fmt.Println(float64(j)*x)

  /*
  // GO demands correct type
  w, z := DoubleAndDuplicate(m)
  fmt.Println(w, z)
  */

  n := 17
  fmt.Println(AddOne(n))
  fmt.Println(n)

  o := SumTwoInts(n, p)
  fmt.Println(o)
}

func AddOne(n int) int { // the variable n gets created as a copy of whatever it is given
  n = n + 1
  return n // return a value
  // the new n is destroyed when the func is completed
}
// SumTwoInts takes 2 ints and returns their sum.
func SumTwoInts(a int, b int) int {
    return a + b
}

func DoubleAndDuplicate(x float64) (float64, float64) {
    return 2.0*x, 2.0*x
}

func Pi() float64 { // doesn't take inputs
  return 3.14 //wrong
}

// functions can also not return anything

func PrintHi() {
  fmt.Println("Hi")
}
