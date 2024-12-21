package main

import (
	"fmt"
	"image"
)

func main() {
	// Create a new image
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(img.Bounds())
	fmt.Println(img.At(0, 0).RGBA())
}

//Bounds() is returning the rectangle of the image and At() is returning the color of the pixel at the given point.
//golang can use the image package to create and manipulate images.
//The image package provides a common interface to image creation and manipulation.
