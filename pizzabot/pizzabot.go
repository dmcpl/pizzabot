package pizzabot

import (
	"github.com/dmcpl/pizzabot/algorithms"
	"github.com/dmcpl/pizzabot/cli"
	"github.com/dmcpl/pizzabot/grid"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"time"
)

// BotRunner this is the main harness for the PizzaBot
// Executes the pathfinding algorithm chosen, creates results and writes output.
func BotRunner(config cli.Config, path grid.Path) {
	pathFindingAlgorithm, ok := algorithms.Algorithms[config.Algorithm]

	if !ok {
		log.Fatalf("Algorithm %s not found!", config.Algorithm)
	}

	t := time.Now()
	paths := pathFindingAlgorithm(path)
	duration := time.Since(t)

	results := createPathResultsMap(paths)

	if config.Verbose {
		WriteVerboseOutput(os.Stdout, results, duration)
	} else {
		WriteBasicOutput(os.Stdout, results)
	}
}

// createPathResultsMap creates a map of Paths keyed by their total distance
func createPathResultsMap(paths []grid.Path) map[int][]grid.Path {
	results := make(map[int][]grid.Path, 0)
	for _, p := range paths {
		results[p.Distance()] = append(results[p.Distance()], p)
	}

	return results
}

// WriteBasicOutput output a single line representing the result path
func WriteBasicOutput(w io.Writer, results map[int][]grid.Path) {
	lengths := make([]int, 0, len(results))
	for k := range results {
		lengths = append(lengths, k)

	}
	sort.Ints(lengths)

	fmt.Fprint(w, results[lengths[0]][0].StringifyPath())
}

// WriteVerboseOutput out put all paths found and time taken to calculate them
func WriteVerboseOutput(w io.Writer, results map[int][]grid.Path, duration time.Duration) {

	fmt.Fprint(w, "Basic path output:\n")
	WriteBasicOutput(w, results)

	lengths := make([]int, 0, len(results))
	for k := range results {
		lengths = append(lengths, k)

	}
	sort.Ints(lengths)

	for _, i := range lengths {
		fmt.Fprintf(w, "\n\nPaths with %d steps:\n", i)
		for _, j := range results[i] {
			fmt.Fprintf(w, "%s\n", j)
			fmt.Fprintf(w, "%s\n", j.StringifyPath())
		}
	}

	fmt.Fprintf(w, "\n\nRun took %dms", duration.Milliseconds())
}
