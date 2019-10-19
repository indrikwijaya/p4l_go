package main

import (
    "fmt"
    "math"
    "time"
    "log"
)

func main() {
    fmt.Println("Primes and arrays.")

    // arrays in Go have a fixed, constant size
    var list [6]int // gives a 6 default values
    // var patterns [28]string // gives 28 "" (empty strings)
    list[0] = -8

    i := 3
    list[2*i-4] = 17
    list[5] = 43 // or list[len(list)-1] = 43 len(list)-1 is the last index
    // list[len(list)] = 5 will give error as the length is only 6
    // list[-4] = 0 also gives error as it doesn't accept negative index
    fmt.Println(list)

    n := 10000

    start := time.Now()
    TrivialPrimeFinder(n)
    elapsed := time.Since(start)
    log.Printf("TrivialPrimeFinder took  %s", elapsed)

    start2 := time.Now()
    SieveofEratosthenes(n)
    elapsed2 := time.Since(start2)
    log.Printf("SieveofEratosthenes took %s", elapsed2)
}


// remember that array has constant size in go.

// we use something called a "slice" to represent variable sizes in Go.
// even if you're plugging in a variable into the length of the array and will
// never change it.

// TrivialPrimeFinder takes an integer n and produces a boolean array of length
// n+1 where primeArray[p] is true if p is prime and false o.w.
func TrivialPrimeFinder(n int) []bool {
    // var primeArray [n+1]bool // default to false, but this initialisation will
                            // lead to error as array shd be fixed
    var primeArray []bool // default to false, this is array with variable length ie slice bool
    // we need an initial length
    // slices start as nil and need to be made
    primeArray = make([]bool, n+1)

    // in general, you'll just use primeArray := make([]bool, n+1)

    for p := 2; p <= n; p++ {
        if IsPrime(p) == true {
            primeArray[p] = true
        }
    }

    return primeArray
}

func IsPrime(p int) bool {
    for k := 2; float64(k) <= math.Sqrt(float64(p)); k++ {
        if p%k == 0 { // p is not prime
            return false
        }
    }
    // if we survive testing all these factors then p is prime
    return true
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
