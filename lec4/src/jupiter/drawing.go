package main

import (
    "image"
    "canvas"
)
// let's place our drawing functions here.

//AnimateSystem takes a collection of Universe objects along with a canvas width
//parameter and generates a slice of images corresponding to drawing each Universe
//on a canvasWidth x canvasWidth canvas
func AnimateSystem(timePoints []Universe, canvasWidth int) []image.Image {
    images := make([]image.Image, len(timePoints))

    for i := range timePoints {
        images[i] = DrawToCanvas(timePoints[i], canvasWidth)
    }
    return images
}

//DrawToCanvas generates the image corresponding to a canvas after drawing a Universe
//object's bodies on a square canvas that is canvasWidth pixels x canvasWidth pixels
func DrawToCanvas(u Universe, canvasWidth int) image.Image {
    // set a new square canvas
    c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

    // create a black background
    c.SetFillColor(MakeColor(0, 0, 0))
    c.ClearRect()
}
