package cli

import (
	"example.com/slice/pizzabot/grid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseGridSize(t *testing.T) {
	argument := "5x5 (0, 0) (1, 3) (4, 4) (4, 2) (4, 2) (0, 1) (3, 2) (2, 3) (4, 1)"
	b, err := parseGridSize(argument)
	assert.Equal(t, &grid.Size{5, 5}, b)
	assert.Nil(t, err)

	argument = " 66 X 77 (0, 0) (1, 3) (4, 4) (4, 2) (4, 2) (0, 1) (3, 2) (2, 3) (4, 1)"
	b, err = parseGridSize(argument)
	assert.Nil(t, err)

	argument = " 66 Xx 77 (0, 0) (1, 3) (4, 4) (4, 2) (4, 2) (0, 1) (3, 2) (2, 3) (4, 1)"
	b, err = parseGridSize(argument)
	assert.Nil(t, b)
	assert.NotNil(t, err)

	argument = " 66s X 77 (0, 0) (1, 3) (4, 4) (4, 2) (4, 2) (0, 1) (3, 2) (2, 3) (4, 1)"
	b, err = parseGridSize(argument)
	assert.Nil(t, b)
	assert.NotNil(t, err)

	argument = "66x77s (0, 0) (1, 3) (4, 4) (4, 2) (4, 2) (0, 1) (3, 2) (2, 3) (4, 1)"
	b, err = parseGridSize(argument)
	assert.Equal(t, &grid.Size{66, 77}, b)
	assert.Nil(t, err)

	argument = "0x5 (0, 0) (1, 3) (4, 4) (4, 2) (4, 2) (0, 1) (3, 2) (2, 3) (4, 1)"
	b, err = parseGridSize(argument)
	assert.Nil(t, b)
	assert.NotNil(t, err)

	argument = "0.7x5 (0, 0) (1, 3) (4, 4) (4, 2) (4, 2) (0, 1) (3, 2) (2, 3) (4, 1)"
	b, err = parseGridSize(argument)
	assert.Equal(t, &grid.Size{7, 5}, b)
	assert.Nil(t, err)

	argument = "(0, 0) (1, 3) (4, 4) (4, 2) 7x5 (4, 2) (0, 1) (3, 2) (2, 3) (4, 1)"
	b, err = parseGridSize(argument)
	assert.Equal(t, &grid.Size{7, 5}, b)
	assert.Nil(t, err)

	argument = "7x5(0, 0) (1, 3) (4, 4) (4, 2)  (4, 2) (0, 1) (3, 2) (2, 3) (4, 1)"
	b, err = parseGridSize(argument)
	assert.Equal(t, &grid.Size{7, 5}, b)
	assert.Nil(t, err)
}

func TestParsePath(t *testing.T) {
	argument := "5x5 (0, 0) (1, 3) (4, 4) (4, 2) (4, 2) (0, 1) (3, 2) (2, 3) (4, 1)"
	p, err := parsePath(argument)

	expected := grid.NewPath(&grid.Point{0, 0},
		&grid.Point{1, 3},
		&grid.Point{4, 4},
		&grid.Point{4, 2},
		&grid.Point{4, 2},
		&grid.Point{0, 1},
		&grid.Point{3, 2},
		&grid.Point{2, 3},
		&grid.Point{4, 1})

	assert.Equal(t, expected, p)
	assert.Nil(t, err)

	argument = "5x5 (0,0) (1,3) (4,4) (4,2) (4,2) (0,1) (3,2) (2,3) (4,1)"
	p, err = parsePath(argument)
	assert.Equal(t, expected, p)
	assert.Nil(t, err)

	argument = "5x5 (0,0) (1,3) (4,4) (4,2) (4,2) (0,1) (3,2) (2,3) (s4s2,1)"
	p, err = parsePath(argument)
	assert.Equal(t, expected.Points[:len(expected.Points)-1], p.Points)
}

func TestValidatePoints(t *testing.T) {
	size := &grid.Size{5, 5}
	path := grid.NewPath(&grid.Point{0, 0},
		&grid.Point{1, 3},
		&grid.Point{4, 4},
		&grid.Point{4, 2},
		&grid.Point{4, 2},
		&grid.Point{0, 1},
		&grid.Point{3, 2},
		&grid.Point{2, 3},
		&grid.Point{4, 1})

	err := validatePoints(path, size)
	assert.Nil(t, err)

	path = grid.NewPath(&grid.Point{0, 0},
		&grid.Point{1, 9},
		&grid.Point{4, 4},
		&grid.Point{10, 2},
		&grid.Point{4, 2},
		&grid.Point{0, 1},
		&grid.Point{5, 2},
		&grid.Point{2, 6},
		&grid.Point{4, 1})

	err = validatePoints(path, size)
	assert.NotNil(t, err)
}
