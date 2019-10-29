package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println("More on arrays/slices.")

    a := make([]int, 3)
    var b [3]int

    ChangeFirst1Slice(a)
    ChangeFirst1Array(b)

    fmt.Println(a)
    fmt.Println(b)

    a = append(a, 5) // this adds element 5 to right of the list
    fmt.Println(a)

    n := 30
    primes := ListPrimes(n)
    fmt.Println(primes)
}

func Max(list []int) int {
    if len(list) == 0 {
        panic("Error: Empty list passed to Max().")
    }

    var m int // default value = 0

    // range over list, updating m if we find a bigger value
    for i, val := range list { // equivalent to for i := 0; i < len(list); i++ {
        if i == 0 || val > m { // val = list[i]
            m = val
        }
    }
    return m
}


// Sum takes a slice of integers and returns the sum of values in the slice
func Sum(a []int) int {
    var s int

    for _, value := range a { // _ says "I don't need this variable"
        s += value
    }

    return s
}

// passing a slice as a parameter

func ChangeFirst1Slice(list []int) {
    list[0] = 1 // we can edit a slice by passing it into a function
}

func ChangeFirst1Array(list [3]int) {
    list[0] = 1 // list is a copy of whatever we call function on
}

//ListPrimes takes an integer n and returns a list of all prime numbers up to and including n.
func ListPrimes(n int) []int {
    primes := make([]int, 0) // great if you don't know the eventual length of primes

    primeBooleans := SieveofEratosthenes(n) // gives me a slice of boolean variables whose p-th element is true if p is prime
    for p := range primeBooleans {
        // when I encounter a prime, append it to primes
        if primeBooleans[p] == true {
            primes = append(primes, p)
        }
    }

    return primes
}

// SieveofEratosthenes takes an integer n and returns a slice of n+1 booleans
// where primeArray[p] is true if p is prime and false o.w.
// It implements the Sieve of Eratosthenes approach.
func SieveofEratosthenes(n int) []bool {
    primeArray := make([]bool, n+1)

    // set everything to prime other than 0 and 1
    for k := 2; k <= n; k++ {
        primeArray[k] = true
    }

    // now range over primeArray, and cross off multiples of first prime we see
    // and iterate this process
    for p := 2; float64(p) <= math.Sqrt(float64(n)); p++ {
        if primeArray[p] == true {
          primeArray = CrossOffMultiples(primeArray, p)
        }
}

    return primeArray
}

// CrossOffMultiples takes a boolean slice and integer p. It crosses off
// multiples of p, meaning turn these multiples to false in the slice
// it returns the slice obtained as a result.
func CrossOffMultiples(primeArray []bool, p int) []bool {
    n := len(primeArray) - 1
    for k := 2*p; k <= n; k += p {
        // all these multiples should be made composite
        primeArray[k] = false
    }

    return primeArray
}
