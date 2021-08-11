package algorithms

import (
	"example.com/slice/pizzabot/grid"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserOrdered(t *testing.T) {
	path := grid.NewPath(grid.Point{1, 0},
		grid.Point{0, 1},
		grid.Point{2, 1},
		grid.Point{1, 2})
	paths := UserOrdered(path)
	expectedPath := grid.NewPath(grid.Point{0, 0},
		grid.Point{1, 0},
		grid.Point{0, 1},
		grid.Point{2, 1},
		grid.Point{1, 2})

	var expectedPaths []grid.Path
	expectedPaths = append(expectedPaths, expectedPath)

	assert.Equal(t, expectedPaths, paths)
}

func TestEmptyOrdered(t *testing.T) {
	path := grid.NewPath()
	paths := UserOrdered(path)
	expectedPath := grid.NewPath(grid.Point{0,0})

	var expectedPaths []grid.Path
	expectedPaths = append(expectedPaths, expectedPath)

	assert.Equal(t, expectedPaths, paths)
}

func TestClosestPoint(t *testing.T) {
	path := grid.NewPath(grid.Point{1, 0},
		grid.Point{0, 1},
		grid.Point{2, 1},
		grid.Point{1, 2})
	paths := ClosestPoint(path)
	expectedPath := grid.NewPath(grid.Point{0, 0},
		grid.Point{1, 0},
		grid.Point{0, 1},
		grid.Point{2, 1},
		grid.Point{1, 2})

	var expectedPaths []grid.Path
	expectedPaths = append(expectedPaths, expectedPath)

	assert.Equal(t, expectedPaths, paths)
}

func TestFindClosestPoint(t *testing.T) {
	a := grid.Point{0, 0}

	p := grid.NewPath(grid.Point{1, 0},
		grid.Point{0, 1},
		grid.Point{2, 1},
		grid.Point{1, 2})

	expected := grid.NewPath(grid.Point{1, 0}, grid.Point{0, 1})

	assert.Equal(t, expected, findClosestPoints(a, p))
}

func TestBuildCPTreeLinear(t *testing.T) {
	origin := grid.Point{0, 0}
	path := grid.NewPath(grid.Point{0, 0},
		grid.Point{0, 1},
		grid.Point{1, 3},
		grid.Point{4, 4},
		grid.Point{4, 2},
		grid.Point{4, 2},
		grid.Point{0, 1},
		grid.Point{3, 2},
		grid.Point{2, 3},
		grid.Point{4, 1})

	node := buildTree(origin, path, findClosestPoints)
	assert.NotNil(t, node)

	acc := make([]grid.Path, 0)
	traverseTree(node, nil, &acc)

	assert.Equal(t, path, grid.NewPath(acc[0].Points[1:]...))

	fmt.Println(path)
	fmt.Println(grid.NewPath(acc[0].Points[1:]...))

	//todo check acc for correct paths
}

func TestBuildCPTreeTwoPaths(t *testing.T) {
	origin := grid.Point{0, 0}
	path := grid.NewPath(grid.Point{0, 0},
		grid.Point{1, 1},
		grid.Point{1, 2},
		grid.Point{2, 1},
		grid.Point{3, 3},
		grid.Point{4, 4})

	node := buildTree(origin, path, findClosestPoints)
	assert.NotNil(t, node)

	acc := make([]grid.Path, 0)
	traverseTree(node, nil, &acc)

	for _, path := range acc {
		fmt.Println(path)
	}

	//todo check acc for correct paths
}

func TestBuildCPTreeMultiPaths(t *testing.T) {
	origin := grid.Point{0, 0}
	path := grid.NewPath(grid.Point{1, 1},
		grid.Point{2, 1},
		grid.Point{2, 1},
		grid.Point{2, 1},
		grid.Point{1, 4},
		grid.Point{0, 4},
		grid.Point{0, 5},
		grid.Point{1, 5},
		grid.Point{2, 5},

		grid.Point{5, 5})

	node := buildTree(origin, path, findClosestPoints)
	assert.NotNil(t, node)

	acc := make([]grid.Path, 0)
	traverseTree(node, nil, &acc)

	for _, path := range acc {
		fmt.Println(path)
	}

	//todo check acc for correct paths
}

func TestBuildBFTree(t *testing.T) {
	origin := grid.Point{0, 0}
	path := grid.NewPath(grid.Point{1, 1},
		grid.Point{2, 1},
		grid.Point{2, 1},
		grid.Point{2, 1},
		grid.Point{1, 4},
	grid.Point{5, 5})

	node := buildTree(origin, path, findAllPoints)
	assert.NotNil(t, node)

	acc := make([]grid.Path, 0)
	traverseTree(node, nil, &acc)

	for _, path := range acc {
		fmt.Println(path)
	}

	assert.Equal(t, 120, len(acc))

	//todo check acc for correct paths
}

func TestAvailableAlgorithms(t *testing.T) {

	t.Fail()
}
