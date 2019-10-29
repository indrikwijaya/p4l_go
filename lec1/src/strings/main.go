package main

import (
    "fmt"
    //"strconv"
    "strings"
)

func main() {
    fmt.Println("Strings.")

    /*
    fmt.Println(string('A'))

    fmt.Println(string(45))
    // Go thinks we want the 45-th symbol according to ASCII chart

    mj := strconv.Itoa(45)
    fmt.Println(mj) // prints 45

    var j int
    j, err := strconv.Atoi("-37")
    // if this goes okay, err = nill
    if err != nil { // problem
        panic("Error: issues converting string to integer.")
    }

    fmt.Println(j)

    x, err2 := strconv.ParseFloat("3.14", 64)
    if err2 != nil { // problem
        panic("Error: issues converting string to float.")
    }
    fmt.Println(x)

    // recall: string concatenation is reprensented by "+"
    s := "Hi"
    t := "Lovers"
    u := s + t
    fmt.Println(u)

    fmt.Println(string(u[0]))
    fmt.Println(string(u[len(u)-1]))

    if t[2] == 'v' {
        fmt.Println("Success!")
    }

    dna := "ACCGAT"
    fmt.Println(Complement(dna))
    fmt.Println(Reverse(dna))
    fmt.Println(ReverseComplement(dna))
    */

    s := "Hi Lovers!"
    fmt.Println(s[2:5]) // substring of s of length 5-2 = 3 starting at position 2.
    fmt.Println(s[4:]) // if I don't specify 2nd index, Go assumes we go to end of strings
    fmt.Println(s[:5]) // prefix substring of length 5

    // note: sub-arrays and subslices have same notation

    a := make([]int, 6)
    for i := range a {
        a[i] = -2*i + 1
    }
    fmt.Println(a)
    fmt.Println(a[3:5])

    // brainteaser: given a slice a, how could we delete an element at a given index?

    index := 2
    a = append(a[:index], a[index+1:]...) //... any time we want to append multiple items.
    fmt.Println(a)

    // replace element at index 1 with element at index 3
    a[1] = a[3]
    fmt.Println(a)
    a = append(a[:3], a[4:]...)
    fmt.Println(a)

    pattern := "ATA"
    text := "ATATA"

    fmt.Println(strings.Count(text, pattern))

}

// ReverseComplement takes a DNA string and returns its reverse complement
// corresponding to opposing strand.

func ReverseComplement(dna string) string {
      return Reverse(Complement(dna))
}

func Reverse(s string) string {
    s2 := ""
    n := len(s)

    for i := range s {
        s2 += string(s[n-1-i])
    }

    return s2
}

func Complement(dna string) string {
    dna2 := ""
    for i := range dna {
        // would give error : can't change symbols of a string in GO
        /*
        switch dna[i] {
        case 'A':
          dna[i] = 'T'
        case 'T':
          dna[i] = 'A'
        case 'C':
          dna[i] = 'G'
        case 'G':
          dna[i] = 'C'
        }
        */
        switch dna[i] {
        case 'A':
          dna2 += "T"
        case 'T':
          dna2 += "A"
        case 'C':
          dna2 += "G"
        case 'G':
          dna2 += "C"
        }
    }
    return dna2
}

// let's count the no of occurrences of a pattern in a text as a substring

func CountPattern(patter, text string) int {
    count := 0
    k := len(pattern)
    n := len(text)

    // range o
}
