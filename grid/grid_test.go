package grid

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbs(t *testing.T) {
	assert.Equal(t, 5, abs(-5))
	assert.Equal(t, 5, abs(5))
	assert.Equal(t, 0, abs(-0))
	assert.Equal(t, 69, abs('E'))
}

func TestDistanceTo(t *testing.T) {
	a := Point{0, 0}
	b := Point{1, 1}

	assert.Equal(t, 2, a.DistanceTo(b))

	a = Point{0, 0}
	b = Point{4, 4}

	assert.Equal(t, 8, a.DistanceTo(b))

	a = Point{4, 1}
	b = Point{1, 5}

	assert.Equal(t, 7, a.DistanceTo(b))
}

func TestRemovePath(t *testing.T) {
	path := NewPath(Point{1, 1}, Point{2, 1}, Point{1, 2}, Point{4, 3})
	newPath := Remove(path, Point{1, 1})
	expected := NewPath(Point{2, 1}, Point{1, 2}, Point{4, 3})

	assert.Equal(t, expected, newPath)

	path = NewPath()
	newPath = Remove(path, Point{1, 1})

	assert.Equal(t, path, newPath)

	path = NewPath(Point{1, 1}, Point{2, 1})
	newPath = Remove(path, Point{1, 1})
	expected = NewPath(Point{2, 1})

	assert.Equal(t, expected, newPath)

	path = NewPath(Point{1, 1}, Point{2, 1})
	newPath = Remove(path, Point{2, 1})
	expected = NewPath(Point{1, 1})

	assert.Equal(t, expected, newPath)

	path = NewPath(Point{1, 1}, Point{2, 1}, Point{3, 3})
	newPath = Remove(path, Point{2, 1})
	expected = NewPath(Point{1, 1}, Point{3, 3})

	assert.Equal(t, expected, newPath)

	path = NewPath(Point{1, 1})
	newPath = Remove(path, Point{1, 1})

	assert.Equal(t, NewPath(), newPath)
	assert.Equal(t, NewPath(Point{1, 1}), path)
}

func TestStringifyPathTo(t *testing.T) {
	a := Point{0, 0}
	b := Point{1, 3}
	output := a.StringifyPathTo(b)

	assert.Equal(t, "ENNN", output)

	a = Point{2, 3}
	b = Point{4, 1}
	output = a.StringifyPathTo(b)

	assert.Equal(t, "EESS", output)

	a = Point{5, 4}
	b = Point{1, 0}
	output = a.StringifyPathTo(b)

	assert.Equal(t, "WWWWSSSS", output)
}

func TestStringifyPath(t *testing.T) {
	path := NewPath(Point{0, 0},
		Point{1, 3},
		Point{4, 4},
		Point{4, 2},
		Point{4, 2},
		Point{0, 1},
		Point{3, 2},
		Point{2, 3},
		Point{4, 1})

	assert.Equal(t, "ENNNDEEENDSSDDWWWWSDEEENDWNDEESSD", path.StringifyPath())

	path = NewPath(Point{5, 0},
		Point{2, 3},
		Point{4, 4},
		Point{5, 2},
		Point{4, 2},
		Point{0, 1},
		Point{1, 2},
		Point{5, 3},
		Point{4, 4})

	assert.Equal(t, "WWWNNNDEENDESSDWDWWWWSDENDEEEENDWND", path.StringifyPath())
}

func TestPathDistance(t *testing.T) {
	path := NewPath(Point{0, 0},
		Point{1, 3},
		Point{4, 4},
		Point{4, 2},
		Point{4, 2},
		Point{0, 1},
		Point{3, 2},
		Point{2, 3},
		Point{4, 1})

	assert.Equal(t, 25,path.Distance())

	path = NewPath(Point{5, 0},
		Point{2, 3},
		Point{4, 4},
		Point{5, 2},
		Point{4, 2},
		Point{0, 1},
		Point{1, 2},
		Point{5, 3},
		Point{4, 4})

	assert.Equal(t, 27,path.Distance())
}

func TestPrintFormatting(t *testing.T) {
	path := NewPath(Point{0, 0},
		Point{1, 3},
		Point{4, 4},
		Point{4, 2},
		Point{4, 2},
		Point{0, 1},
		Point{3, 2},
		Point{2, 3},
		Point{4, 1})

	assert.Equal(t, "[(0, 0) (1, 3) (4, 4) (4, 2) (4, 2) (0, 1) (3, 2) (2, 3) (4, 1)]", fmt.Sprintf("%v", path))

	path = NewPath(Point{5, 0},
		Point{2, 3},
		Point{4, 4},
		Point{5, 2},
		Point{4, 2},
		Point{0, 1},
		Point{1, 2},
		Point{5, 3},
		Point{4, 4})

	assert.Equal(t, "[(5, 0) (2, 3) (4, 4) (5, 2) (4, 2) (0, 1) (1, 2) (5, 3) (4, 4)]", fmt.Sprintf("%v", path))
}

func TestPointOutSideRangeOf(t *testing.T) {
	grid := Size{5, 5}

	point := Point{0, 0}
	assert.False(t, point.OutSideRangeOf(grid))

	point = Point{6, 1}
	assert.True(t, point.OutSideRangeOf(grid))

	point = Point{1, 6}
	assert.True(t, point.OutSideRangeOf(grid))

	point = Point{5, 5}
	assert.False(t, point.OutSideRangeOf(grid))
}