package main

import (
    "fmt"
    "math"
)

type Tree struct {
    nodes []Node
}

type Tree []Node


type Node struct {
    age            float64
    label          string
    child1, child2 *Node
    // by making the children pointers to Node objects, we avoid recursive type!
    // with var x Node, x.age=0, x.label = "",
    // x.child1 = x.child2 = nil
}


type Rectangle struct {
    x1, y1 float64
    width, height float64
    rotation float64
}

type Circle struct {
    x1, y1 float64 // center point
    radius float64
}

func (r Rectangle) Area() float64 {
    return r.width*r.height
}

func (c Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}

func (r Rectangle) Translate(a, b float64) Rectangle {
    // add a to x-coordinate of corner pt
    // add b to y-coordinate of corner pt
    r.x1 += a
    r.y1 += b
    return r
}

func (c Circle) Translate(a, b float64) Circle {
    c.x1 += a
    c.y1 += b
    return c
}

func (r *Rectangle) Translate2(a, b float64) {
    r.x1 += a
    r.y1 += b
}

func (c *Circle) Translate2(a, b float64) {
    c.x1 += a
    c.y1 += b
}

func main() {
    fmt.Println("More object-oriented programming in Go!")
    /*
    var r Rectangle
    var c Circle

    r.width = 3.0
    r.height = 5.0

    c.radius = 2.0

    fmt.Println(r.Area())
    fmt.Println(c.Area())

    // let's translate the rectangle and circle
    fmt.Println(r.x1, r.y1)
    fmt.Println(c.x1, c.y1)

    r = r.Translate(1, -2)
    c = c.Translate(-3, 16)

    fmt.Println(r.x1, r.y1)
    fmt.Println(c.x1, c.y1)
    */

    /*
    // part 2: discussion of pointers
    // pointer: a variable that holds the memory address of some other variable

    var b int = -47

    var a *int // a is a pointer (i.e., address of) to an integer variable
    fmt.Println(b)
    fmt.Println(a) // a currently has default special value "nil"

    // where is b stored in RAM?
    a = &b // & means "location of"
    fmt.Println(a)

    // is there a one-line a := &b? Yes!
    var c Circle
    var pointer *Circle // has value nil

    pointer = &c // now pointer points to c

    // if we have just a pointer to an object, can we change it?
    // Yes.
    // say we want to change radius of c by accessing the pointer we created.
    //* also allows 'dereference' of a pointer.
    //it means 'grab the variable at the address that the pointer refers to.'
    (*pointer).radius = 4.6
    fmt.Println(c.radius)

    // good news: Go will automatically interpret a pointer to an object as the object.
    pointer.radius = 2.33 // this is OK!
    fmt.Println(c.radius)

    c.Translate2(4.3, -0.71)
    //Go will understand to access c, not a copy of it.
    fmt.Println(c.x1, c.y1)
    */

    // part 3: a bit about slices
    a := make([]int, 5) // 0, 0, 0, 0, 0
    ChangeFirst(a)
    fmt.Println(a)

    var b [5]int
    ChangeFirstArray(b) // a copy of b is created!
    fmt.Println(b)

    c := make([]int, 5, 10) // c is a slice of length 5 that points to the first 5 elements of some array of length 10

    for i := range c {
        c[i] = 2*i - 3
        // this sets first five elements of the array. remaining five are 0.
    }

    var d []int = c[3:7] // this is OK!

    fmt.Println("c is", c)
    fmt.Println("d is", d)

    // c and d both point to same arary.
    // so it's possible to update one and it changes the other.

    d[0] = 1820
    fmt.Println("c is", c)
    fmt.Println("d is", d)

    // c[6] = 12

    // let's throw out c[3]
    index := 3
    c = Delete(c, index)
    // c is still pointing at the same place with length 5
    fmt.Println("c is", c)
    fmt.Println("d is", d)

    // brainteaser: how do i get rid of the additional 5 that is contained within d?

}

// A SLICE IS A POINTER TO AN ARRAY (along with a length)

func ChangeFirst(list []int) {
    list[0] = 1
}

func ChangeFirstArray(list [5]int) {
    list[0] = 1
    // copy (list) is destroyed here
}

// Delete removes the element of a given slice of integers at a given index.
func Delete(list []int, index int) []int {
    list = append(list[:index], list[index+1:]...)
    // list is the thing that changes/
    // it points at the same place, it removes the element at the given index, and it gets updated to have length 4.

    // list gets destroyed
    return list // updates the slice outside of the function to have the appropriate length
}
