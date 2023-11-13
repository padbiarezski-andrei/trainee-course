package main

import (
	"flag"
	"fmt"
	"os"
	"trainee-course/interfaces/shapes"
)

func main() {

	var shapeType = flag.String("shape", "", "type of the shape {circle, rectangle}")

	var radius = flag.Float64("radius", -1.0, "circle radius")

	var width = flag.Float64("width", -1.0, "rectangle width")
	var height = flag.Float64("height", -1.0, "rectangle height")

	flag.Parse()

	var shape shapes.Shape

	switch *shapeType {
	case "circle":
		if *radius <= 0.0 {
			fmt.Println("the radius of the circle must be greater than zero")
			os.Exit(1)
		}
		shape = shapes.Circle{R: *radius}
	case "rectangle":
		if *width <= 0.0 || *height <= 0.0 {
			fmt.Println("the side of the rectangle must be greater than zero")
			os.Exit(1)
		}
		shape = shapes.Rectangle{Width: *width, Height: *height}
	default:
		fmt.Println("unknown shape type. Please use --help to see more information about available options")
		os.Exit(1)
	}

	fmt.Println(shape.Area())
}
