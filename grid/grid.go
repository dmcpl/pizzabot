package grid

import (
	"fmt"
	"strings"
)

const (
	East  = "E"
	South = "S"
	North = "N"
	West  = "W"
	Drop  = "D"
)

type Size struct {
	MaxX, MaxY int
}

type Point struct {
	X, Y int
}

func (p Point) DistanceTo(b *Point) int {
	return abs(p.X-b.X) + abs(p.Y-b.Y)
}

func (p Point) StringifyPathTo(b *Point) string {
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

func (p Point) OutSideRangeOf(b *Size) bool {
	if p.X > b.MaxX || p.Y > b.MaxY {
		return true
	}
	return false
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

type Path struct {
	Points []*Point
}

func NewPath(points ...*Point) *Path {
	if len(points) == 0 {
		return &Path{Points: make([]*Point, 0)}
	}
	return &Path{Points: points}
}

// Remove removes an element from the list, returns new path
//todo update this fn, not very nice yet
func (p *Path) Remove(pt *Point) *Path {
	var removalIndex int
	for idx, po := range p.Points {
		if *po == *pt {
			removalIndex = idx
			break
		}
	}

	if removalIndex == 0 {
		return NewPath(p.Points[1:]...)
	} else if removalIndex == len(p.Points)-1 {
		return NewPath(p.Points[:len(p.Points)-1]...)
	} else {
		pre := p.Points[:removalIndex]
		post := p.Points[removalIndex+1:]

		return &Path{append(pre, post...)}
	}
}

func (p *Path) Distance() (distance int) {
	for i := 0; i < len(p.Points)-1; i++ {
		distance += p.Points[i].DistanceTo(p.Points[i+1])
	}
	return distance
}

func (p *Path) StringifyPath() string {
	var b strings.Builder
	for i := 0; i < len(p.Points)-1; i++ {
		b.WriteString(p.Points[i].StringifyPathTo(p.Points[i+1]))
		b.WriteString(Drop)
	}
	return b.String()
}

func (p *Path) String() string {
	return fmt.Sprint(p.Points)
}

// abs Utility function returning the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
