package main

import (
    "fmt"
    "strconv"
)

func main() {
    fmt.Println("Strings.")

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
}

// ReverseComplement takes a DNA string and returns its reverse complement
// corresponding to opposing strand.
func ReverseComplement(dna string) {
      return Reverse(Complement(dna))
}

func Reverse(s string) string {
    // write later
    return s
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
        case "A":
          dna2 += "T"
        case "T":
          dna2 += "A"
        case "C":
          dna2 += "G"
        case "G":
          dna2 += "C"
        }
    }
    return dna2
}
