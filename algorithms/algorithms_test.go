package algorithms

import (
	"example.com/slice/pizzabot/grid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindClosestPoint(t *testing.T) {
	a := &grid.Point{0, 0}

	p := grid.NewPath(&grid.Point{1, 0},
		&grid.Point{0, 1},
		&grid.Point{2, 1},
		&grid.Point{1, 2})

	expected := grid.NewPath(&grid.Point{1, 0}, &grid.Point{0, 1})

	assert.Equal(t, expected, findClosestPoints(a, p))
}

func TestBuildTree(t *testing.T) {
	//todo update test
	origin := &grid.Point{0, 0}
	path := grid.NewPath(&grid.Point{1, 1}, &grid.Point{2, 1}, &grid.Point{1, 2}, &grid.Point{4, 3})

	node := buildTree(origin, path)
	assert.NotNil(t, node)

	acc := make([]*grid.Path, 0)
	traverseTree(node, nil, &acc)

	//todo check acc for correct paths
}
