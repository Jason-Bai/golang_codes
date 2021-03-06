package main

import (
	"fmt"
	"log"

	"github.com/Jason-Bai/golang_codes/packages/geometry/rectangle"
)

var rectLen, rectWidth float64 = 6, 7

// var rectLen, rectWidth float64 = -6, 7

func init() {
	fmt.Println("main package initialized")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func main() {
	fmt.Println("Geometrical shape properties")
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	fmt.Printf("diagonal of the rectangle %.2f\n", rectangle.Diagonal(rectLen, rectWidth))
}
