package main

import "fmt"

func main() {
    fmt.Println("Frequent words and maps.")

    // one eg  of a map: keys are strings(words), values are no of times a
    // given word occurs in a text.
    // another eg: a map whose keys are state names and whose values are polling
    // % (0 to 1) for each candidate.

    /*
    var a []int
    a = make([]int, 6)
    fmt.Println(a)
    */

    // maps generalize the slice declaration
    var polls map[string]float64
    // just like slices, maps need to be made
    polls = make(map[string]float64) // no need to specify length

    // Of course there is a shortcut declaration polls:=
    // make(map[string]float64)
    polls["PA"] = 0.517
    polls["NY"] = 0.789 // no need to append
    polls["NE"] = 0.401
    polls["FL"] = 0.500

    delete(polls, "FL")

    for state, percentage := range polls {
        fmt.Println("The polling percentage of", state, "is", percentage)
    }

}
