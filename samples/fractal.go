package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/julianknodt/turtle"
)

const maxMove = 400.0

func fractal(t turtle.TurtleDrawer, iteration, max int) turtle.TurtleDrawer {
	fmt.Println(iteration)
	if iteration == max {
		return t
	}

	t.Move(maxMove / float64(iteration*iteration+1))

	return fractal(t.Partition(
		turtle.DegreesToRadians(45)),
		iteration+1, max)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	t := &turtle.Turtle{
		X:     400.0,
		XDir:  0,
		YDir:  1.0,
		Color: color.Black,
	}

	t.NewImage(800, 800)

	fractal(t, 0, 12).Save("./fractal.png")
}
