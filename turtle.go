package turtle

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

// TurtleDrawer represents a turtle that scoots around the page
// and has a pen for drawing
type TurtleDrawer interface {
  // Move some distance with the pen on the page
	Move(float64)

  // Rotate the turtle some radians
	Rotate(float64)

  // Partition one turtle into many turtles
	Partition(float64) *TurtleFleet

  // Len is how many turtles are in this drawer
	Len() int

  // Save turtle image into a file
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

// Create a new image for this turtle to draw on
func (t *Turtle) NewImage(height, width int) {
	t.Image = image.NewRGBA(image.Rect(0, 0, width, height))
}

// Len of one turtle is always one
func (t Turtle) Len() int {
	return 1
}

// Move this turtle straight ahead some distance
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

// Rotate this turtle some radians counterclockwise
func (t *Turtle) Rotate(theta float64) {
	cosTheta := math.Cos(theta)
	sinTheta := math.Sin(theta)
	currXDir := t.XDir
	t.XDir = t.XDir*cosTheta - t.YDir*sinTheta
	t.YDir = currXDir*sinTheta + t.YDir*cosTheta
}

// Make a new turtle through the power of science
func (t *Turtle) Clone() *Turtle {
	result := new(Turtle)
	*result = *t
	return result
}

// Make one turtle into two and call it a fleet
//
// The two get rotated by +theta and -theta
func (t *Turtle) Partition(theta float64) *TurtleFleet {
	clone1 := t.Clone()
	clone2 := t.Clone()
	clone1.Rotate(theta)
	clone2.Rotate(-theta)
	return &TurtleFleet{
		turtles: []*Turtle{clone1, clone2},
	}
}

// Convenient function to turn degrees into rads
func DegreesToRadians(deg float64) float64 {
	return deg * math.Pi / 180
}

// Save your turtle to a file
func (t Turtle) Save(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	return png.Encode(f, t.Image)
}
