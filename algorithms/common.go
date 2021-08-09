package algorithms

import (
	"example.com/slice/pizzabot/grid"
)

type AlgorithmFunc func(*grid.Path) []*grid.Path

var Algorithms map[string]AlgorithmFunc

func init() {
	Algorithms = map[string]AlgorithmFunc{"treeCP": NodeTreeClosestPoint,
		"treeBF":  NodeTreeBruteForce,
		"CP":      ClosestPoint,
		"ordered": UserOrdered}
}

func AvailableAlgorithms() (a []string) {
	for k, _ := range Algorithms {
		a = append(a, k)
	}
	return a
}

// FindClosestPoints takes an origin point and calculates the distance to all other points in a path
// returning the closest point(s)
func findClosestPoints(origin *grid.Point, points *grid.Path) *grid.Path {
	var closestDistance int
	closestPoints := grid.NewPath()
	for _, p := range points.Points {
		if d := origin.DistanceTo(p); len(closestPoints.Points) == 0 || d < closestDistance {
			closestPoints.Points = []*grid.Point{p}
			closestDistance = d
		} else if origin.DistanceTo(p) == closestDistance {
			closestPoints.Points = append(closestPoints.Points, p)
		}
	}
	return closestPoints
}
