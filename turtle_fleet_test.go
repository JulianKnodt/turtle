package turtle

import (
	"fmt"
	"testing"
)

func TestTurtleFleet(t *testing.T) {
	turt := &Turtle{}
	fmt.Println(turt.Partition(0).Partition(0).Partition(0).Len())
}
