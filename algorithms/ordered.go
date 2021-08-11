package algorithms

import (
	"example.com/slice/pizzabot/grid"
)

// UserOrdered just returns the path which the user has specified in the parameters prepended by the origin point.
// This will only ever produce one result.
func UserOrdered(path grid.Path) []grid.Path {
	origin := grid.Point{0, 0}
	p := grid.NewPath(origin)
	p.Points = append(p.Points, path.Points...)

	paths := make([]grid.Path, 0, 1)
	paths = append(paths, p)
	return paths
}
