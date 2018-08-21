package main

import (
	"image/color"
	"math"
	"time"

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

	go func() {
		i := 0.0
		for {
			t.Move(math.Log(i))
			t.Rotate(turtle.DegreesToRadians(3))
			i++
		}
	}()

	time.Sleep(3 * time.Second)
	t.Save("./spiral.png")
}
