package turtle

type TurtleFleet struct {
	turtles []*Turtle
}

func (t TurtleFleet) Move(dist float64) {
	for _, t := range t.turtles {
		t.Move(dist)
	}
}

func (t TurtleFleet) Rotate(theta float64) {
	for _, t := range t.turtles {
		t.Rotate(theta)
	}
}

func (t *TurtleFleet) Combine(other TurtleFleet) *TurtleFleet {
	t.turtles = append(t.turtles, other.turtles...)
	return t
}

func (t *TurtleFleet) Partition(theta float64) *TurtleFleet {
	result := &TurtleFleet{turtles: []*Turtle{}}
	for _, turt := range t.turtles {
		result.Combine(*turt.Partition(theta))
	}
	return result
}

func (t TurtleFleet) Len() int {
	return len(t.turtles)
}

func (t TurtleFleet) Save(filename string) error {
	return t.turtles[0].Save(filename)
}
