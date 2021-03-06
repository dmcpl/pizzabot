package cli

import (
	"github.com/dmcpl/pizzabot/grid"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// parseGridSize extracts a grid size in the form: "int ('x' OR 'X') int"
// Whitespace is allowed around and between the characters for a bit of user-friendliness.
// First match on the argument line is used, others ignored.
func parseGridSize(argument string) (gridSize grid.Size, err error) {
	gridSizePattern := "\\s*\\d+\\s*(X|x)\\s*\\d+\\s*"
	matcher, err := regexp.Compile(gridSizePattern)
	if err != nil {
		return gridSize, fmt.Errorf("cannot parse grid pattern: %v", err)
	}

	match := matcher.FindString(argument)
	if len(match) > 0 { // Found a potential grid size match
		numbersOnly := strings.NewReplacer("X", " ", "x", " ").Replace(match) // replace the 'X' with whitespace
		fields := strings.Fields(numbersOnly)                                 // extract all non whitespace strings

		if len(fields) != 2 {
			return gridSize, fmt.Errorf("cannot extract grid size from fields %v", fields)
		}

		x, err := strconv.Atoi(fields[0])
		if err != nil || x <= 0 {
			return gridSize, fmt.Errorf("cannot parse horizontal grid size from argument (minimum value is 1): %s", fields[0])
		}
		y, err := strconv.Atoi(fields[1])
		if err != nil || y <= 0 {
			return gridSize, fmt.Errorf("cannot parse vertical grid size from argument (minimum value is 1): %s", fields[1])
		}
		gridSize = grid.Size{x, y}
	} else {
		return gridSize, fmt.Errorf("cannot parse grid size from argument: %v", argument)
	}

	return gridSize, nil
}

// parsePath extracts a list in the forms "(10,5) (5,17)..."
// Whitespace is allowed around and between the characters for a bit of user-friendliness.
// All matches to this pattern are accepted, everything else is ignored.
func parsePath(argument string) (path grid.Path, err error) {
	pointPattern := "\\s*\\(\\s*\\d+\\s*,\\s*\\d+\\s*\\)\\s*"
	matcher, err := regexp.Compile(pointPattern)
	if err != nil {
		return path, fmt.Errorf("cannot parse grid pattern: %v", err)
	}

	points := make([]grid.Point, 0)
	rawPoints := matcher.FindAllString(argument, -1)
	for _, rawPoint := range rawPoints {
		numbersOnly := strings.NewReplacer("(", " ", ")", " ", ",", " ").Replace(rawPoint) // replace  ')' '(' and ',' with whitespace
		fields := strings.Fields(numbersOnly)

		if len(fields) != 2 {
			return path, fmt.Errorf("cannot extract point coordinates from fields %v", fields)
		}

		x, err := strconv.Atoi(fields[0])
		if err != nil || x < 0 {
			return path, fmt.Errorf("cannot parse X coodinate from argument (minimum value is 0): %s", rawPoint)
		}
		y, err := strconv.Atoi(fields[1])
		if err != nil || y < 0 {
			return path, fmt.Errorf("cannot parse Y coodinate from argument (minimum value is 0): %s", rawPoint)
		}
		points = append(points, grid.Point{x, y})
	}

	return grid.NewPath(points...), nil
}

// validatePoints returns an error if any of the path's points are outside the grid size
func validatePoints(path grid.Path, gridSize grid.Size) error {
	invalidPath := grid.NewPath()
	for _, p := range path.Points {
		if p.OutSideRangeOf(gridSize) {
			invalidPath.Points = append(invalidPath.Points, p)
		}
	}
	if len(invalidPath.Points) > 0 {
		return fmt.Errorf("invalid points encountered: %v", invalidPath)
	}

	return nil
}

// ParseBotParam parses the arguments used to specify grid size and the various drop points
func ParseBotParam(argument string) (path grid.Path, err error) {

	gridSize, err := parseGridSize(argument)
	if err != nil {
		return path, err
	}

	path, err = parsePath(argument)
	if err != nil {
		return path, err
	}

	err = validatePoints(path, gridSize)
	if err != nil {
		return path, err
	}

	return path, nil
}
