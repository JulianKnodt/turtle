package turtle

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

type TurtleDrawer interface {
	Move(float64)
	Rotate(float64)
	Partition(float64) *TurtleFleet
	Len() int
	Save(filename string) error
}

// Turtle draws for every unit he moves
type Turtle struct {
	X     float64
	Y     float64
	XDir  float64
	YDir  float64
	Color color.Color
	PenUp bool

	Image *image.RGBA
}

func (t *Turtle) NewImage(height, width int) {
	t.Image = image.NewRGBA(image.Rect(0, 0, width, height))
}

func (t Turtle) Len() int {
	return 1
}

func (t *Turtle) Move(dist float64) {
	for i := 0.0; i < dist; i++ {
		t.X += t.XDir
		t.Y += t.YDir
		if t.PenUp {
			continue
		}
		x, y := int(t.X), int(t.Y)
		if t.Image != nil && image.Pt(x, y).In(t.Image.Bounds()) {
			t.Image.Set(x, y, t.Color)
		}
	}
}

func (t *Turtle) Rotate(theta float64) {
	cosTheta := math.Cos(theta)
	sinTheta := math.Sin(theta)
	currXDir := t.XDir
	t.XDir = t.XDir*cosTheta - t.YDir*sinTheta
	t.YDir = currXDir*sinTheta + t.YDir*cosTheta
}

func (t *Turtle) Clone() *Turtle {
	result := new(Turtle)
	*result = *t
	return result
}

func (t *Turtle) Partition(theta float64) *TurtleFleet {
	clone1 := t.Clone()
	clone2 := t.Clone()
	clone1.Rotate(theta)
	clone2.Rotate(-theta)
	return &TurtleFleet{
		turtles: []*Turtle{clone1, clone2},
	}
}

func DegreesToRadians(deg float64) float64 {
	return deg * math.Pi / 180
}

func (t Turtle) Save(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	return png.Encode(f, t.Image)
}
