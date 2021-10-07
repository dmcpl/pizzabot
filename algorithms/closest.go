package algorithms

import (
	"github.com/dmcpl/pizzabot/grid"
)

// ClosestPoint will calculate the path around the grid finding the first closest point to the current point and
// following it. If there are more than one closest point this will choose the first iterated over in the list and
// ignore the others. This will only ever produce one result.
func ClosestPoint(path grid.Path) []grid.Path {
	origin := grid.Point{0, 0}

	result := grid.NewPath(origin)

	for len(path.Points) > 0 {
		cp := findClosestPoints(origin, path)
		grid.Remove(cp, cp.Points[0])
		result.Points = append(result.Points, cp.Points[0])
		path = grid.Remove(path, cp.Points[0])
	}

	results := make([]grid.Path, 0)
	return append(results, result)
}
