package algorithms

import (
	"github.com/dmcpl/pizzabot/grid"
)

type AlgorithmFunc func(grid.Path) []grid.Path

var Algorithms map[string]AlgorithmFunc

func init() {
	Algorithms = map[string]AlgorithmFunc{
		"CP":      ClosestPoint,
		"ordered": UserOrdered,
		"treeCP":  NodeTreeClosestPoint,
		"treeBF":  NodeTreeBruteForce}
}

func AvailableAlgorithms() (a []string) {
	for k, _ := range Algorithms {
		a = append(a, k)
	}
	return a
}

// findClosestPoints takes an origin point and calculates the distance to all other points in a path
// returning the closest point(s)
func findClosestPoints(origin grid.Point, points grid.Path) grid.Path {
	var closestDistance int
	closestPoints := grid.NewPath()
	for _, p := range points.Points {
		if d := origin.DistanceTo(p); len(closestPoints.Points) == 0 || d < closestDistance {
			closestPoints.Points = []grid.Point{p}
			closestDistance = d
		} else if origin.DistanceTo(p) == closestDistance {
			closestPoints.Points = append(closestPoints.Points, p)
		}
	}
	return removeDuplicatePoints(closestPoints)
}

// removeDuplicatePoints Removes any duplicate points from the closest Path, this happens when the user
// specifies points more than once. We only need to hit it once and drop multiple times.
func removeDuplicatePoints(path grid.Path) grid.Path {
	pm := make(map[grid.Point]bool)
	uniquePoints := make([]grid.Point, 0)

	for _, p := range path.Points {
		if value := pm[p]; !value {
			pm[p] = true
			uniquePoints = append(uniquePoints, p)
		}
	}

	return grid.NewPath(uniquePoints...)
}

func findAllPoints(origin grid.Point, points grid.Path) grid.Path {
	return points
}
