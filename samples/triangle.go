package main

import (
	"image/color"

	"github.com/julianknodt/turtle"
)

func main() {

	t := turtle.Turtle{
		X:     400.0,
		Y:     400.0,
		XDir:  1,
		YDir:  0,
		Color: color.Black,
	}

	t.NewImage(800, 800)

	t.Move(100)
	t.Rotate(turtle.DegreesToRadians(240))
	t.Move(100)
	t.Rotate(turtle.DegreesToRadians(-120))
	t.Move(100)

	t.Save("./triangle.png")
}
