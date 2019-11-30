package main

import (
    "fmt"
    "math/rand"
    "time"
  )

func main() {
    fmt.Println("Craps!")

    // rand.Seed(1)
    // default value of seed is 1 if rand.Seed() is not called.
    // instead: seed off time
    rand.Seed(time.Now().UnixNano())
    /*
    for i := 0; i < 7; i++ {
        fmt.Println(RollDie())
    }
    */
    numTrials := 10000000
    fmt.Println(ComputeHouseEdge(numTrials))
}

func RollDie() int {
    return rand.Intn(6) + 1 // this generate random int 0<=x<=5 similar to (rand.Int()%6) + 1
}

// SumTwoDice takes no input and returns the sum of two simulated six-sided dice.
func SumTwoDice() int {
    return RollDie() + RollDie()
}

// math/rand has 3 built-in functions we will use a lot:
// 1. rand.Int: pseudorandom integer
// 2. rand.Float64: pseudorandom decimal
// 3. rand.Intn: pseudorandom integer btw 0 and n-1, inclusively

// PlayCrapsOnce takes no input parameters and returns true or false dependeing on outcome of a single simulated game of craps.
func PlayCrapsOnce() bool {
    firstRoll := SumTwoDice()
    if firstRoll == 7 || firstRoll == 11{
        return true // winner!
    } else if firstRoll == 2 || firstRoll == 3 || firstRoll == 12 {
        return false // loser!
    } else { // roll again until we hit a 7 or our original roll
        for true {
            newRoll := SumTwoDice()
            if newRoll == firstRoll {
                // winner! :)
                return true
            } else if newRoll == 7 {
                // loser! :()
                return false
            }
        }
    }
    // Go often likes default values at end of function
    panic("We shouldn't be here.")
    return false
}

// ComputeHouseEdge takes an integer numTrials and returns an estimate of the house edge of craps (or whatever binary game) over numTrials simulated games.
func ComputeHouseEdge(numTrials int) float64 {
    // we use count to keep track of money won/lost
    count := 0
    for i := 0; i < numTrials; i++ {
        var outcome bool
        outcome = PlayCrapsOnce()
        // did we win or lose?
        if outcome == true { // win
            count++
        } else {
            // lost
            count--
        }
    }
    // we want to return the average won/lost
    return float64(count)/float64(numTrials)
}
