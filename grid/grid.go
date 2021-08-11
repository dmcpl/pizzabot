package grid

import (
	"fmt"
	"log"
	"strings"
)

const (
	East  = "E"
	South = "S"
	North = "N"
	West  = "W"
	Drop  = "D"
)

// Size holds the parsed size of the grid
type Size struct {
	MaxX, MaxY int
}

// Point represents a point on the grid
type Point struct {
	X, Y int
}

// DistanceTo calculates the number of steps needed to travel from this point to a specified Point b
func (p Point) DistanceTo(b Point) int {
	return abs(p.X-b.X) + abs(p.Y-b.Y)
}

// StringifyPathTo returns the steps travelled between two points as a string
// E is for a step East, W for West, N for North, S for south, and D for a Drop at the destination Point
// This particular technique always writes horizontal moves first
func (p Point) StringifyPathTo(b Point) string {
	var s strings.Builder

	diffX := b.X - p.X
	if diffX >= 0 {
		s.WriteString(strings.Repeat(East, abs(diffX)))
	} else {
		s.WriteString(strings.Repeat(West, abs(diffX)))
	}

	diffY := b.Y - p.Y
	if diffY >= 0 {
		s.WriteString(strings.Repeat(North, abs(diffY)))
	} else {
		s.WriteString(strings.Repeat(South, abs(diffY)))
	}

	return s.String()
}

// OutSideRangeOf returns true if the point is inside the grid size, false otherwise
// Does not check minus values.
func (p Point) OutSideRangeOf(b Size) bool {
	if p.X > b.MaxX || p.Y > b.MaxY {
		return true
	}
	return false
}

// String Stringer implementation for printing
func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

// Path represents a list of Points
type Path struct {
	Points []Point
}

// NewPath constructor for Path types, variadic accepts Points.
func NewPath(points ...Point) Path {
	if len(points) == 0 {
		return Path{Points: make([]Point, 0)}
	}
	return Path{Points: points}
}

// Remove removes an element from the list, returns a new path, does not modify the original
func Remove(p Path, pt Point) Path {
	var removalIndex int
	var found bool
	for i, po := range p.Points { // Find the index of the element to remove
		if po == pt {
			removalIndex = i
			found = true
			break
		}
	}

	cp := make([]Point, len(p.Points))
	copied := copy(cp, p.Points) //todo check for written
	if copied != len(p.Points) {
		log.Fatalf("FUCK")
	}

	if found {
		switch {
		case len(p.Points) == 1: // Only one element to remove so empty path
			return NewPath()

		case removalIndex == 0: // Remove head
			return NewPath(cp[1:]...)

		case removalIndex == len(cp)-1: // Remove last
			return NewPath(cp[:len(p.Points)-1]...)

		default: // Somewhere in the middle, join slices before and after the element
			pre := cp[:removalIndex]
			post := cp[removalIndex+1:]
			return NewPath(append(pre, post...)...)
		}
	}

	return NewPath(cp...) // Only if the element is not found, return original
}

// Distance returns the distance in steps around this Path
func (p Path) Distance() (distance int) {
	for i := 0; i < len(p.Points)-1; i++ {
		distance += p.Points[i].DistanceTo(p.Points[i+1])
	}
	return distance
}

// StringifyPath returns the steps travelled around a Path as a string
// E is for a step East, W for West, N for North, S for south, and D for a Drop at the destination Point
// This particular technique always writes horizontal moves first
func (p Path) StringifyPath() string {
	var b strings.Builder
	for i := 0; i < len(p.Points)-1; i++ {
		b.WriteString(p.Points[i].StringifyPathTo(p.Points[i+1]))
		b.WriteString(Drop)
	}
	return b.String()
}

// String Stringer implementation for printing
func (p Path) String() string {
	return fmt.Sprint(p.Points)
}

// abs Utility function returning the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
