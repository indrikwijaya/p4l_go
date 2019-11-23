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

    // array and slice literals
    b := [4]float64{3.1, 0.0, -78.93, 21.22}
    primes := []int{2, 3, 5, 7, 11}

    // map literals too:
    electoralVotes := map[string]uint {
        "PA":20,
        "OH":18,
        "TX":38, // Go demands this final comma in case I add more later
    }
    fmt.Println(b, primes, electoralVotes)

    dict2 := make(map[string]int)
    dict2["Love"] = 0
    BoostLove(dict2)
    fmt.Println(dict2["Love"])

    text := "ACGTTTTGAGACGTTT"
    k := 3
    fmt.Println(FindFrequentWords(text, k))
}

func FindFrequentWords(text string, k int) []string {
    freqPatterns := make([]string, 0)
    freqMap := FrequencyMap(text, k) // will give us a map of k-mers in text to
    // to their no of occurrences
    max := MaxValue(freqMap) // grab max vale in frequency map

    // range over frequency map, looking for strings (keys) achieving the max
    // no of values
    for pattern, val := range freqMap {
        if val == max {
            // append pattern to freqPatterns list
            freqPatterns = append(freqPatterns, pattern)
        }
    }

    return freqPatterns
}

// MaxValue takes a map of strings to ints as input and return the max value in the map.
func MaxValue(freqMap map[string]int) int {
    m := 0
    firstTimeThrough := true

    for _, value := range freqMap {
        if firstTimeThrough || value > m {
            m = value
            firstTimeThrough = false
        }
    }

    return m
}

//FrequencyMap takes a string text and an integer k,
//It returns the map of each k-mer substring of text to its member of
//occurences in text, including overlaps.
func FrequencyMap(text string, k int) map[string]int {
    freq := make(map[string]int)
    // range over all k-mer substrings of text
    for j :=0; j < len(text)-k+1; j++ {
        pattern := text[j:j+k]

        //method 1:
        /*
        // does freq[pattern] exist?
        _, exists := freq[pattern]
        // exists is boolean = false if freq[pattern] doesn't exist
        if !exists {
            freq[pattern] = 1
        } else {
            freq[pattern]++
        }
        */

        //method 2: shortcut
        freq[pattern]++
        // if freq[pattern] doesnt exist, Go creates it as default value = 0
        // and then adds 1 to it. If it exists, it adds 1 to it
    }

    return freq
}
// lesson: like slices, we have access to maps if we pass them into functions
// as parameters.

func BoostLove(dict map[string]int) {
    dict["Love"] = 100
}
