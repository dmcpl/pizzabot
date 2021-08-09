package pizzabot

import (
	"example.com/slice/pizzabot/algorithms"
	"example.com/slice/pizzabot/cli"
	"example.com/slice/pizzabot/grid"
	"fmt"
	"io"
	"os"
	"time"
)

func WriteVerboseOutput(w io.Writer, results map[int][]*grid.Path, duration time.Duration) {
	WriteBasicOutput(w, results)

	//todo write all paths
	//todo write duration taken
}

func WriteBasicOutput(w io.Writer, results map[int][]*grid.Path) {
	for n, paths := range results {
		fmt.Fprintf(w, "Results with %d moves:\n", n)
		for _, p := range paths {
			fmt.Fprintf(w, "%s\n", p.StringifyPath())
		}
	}
}

// BotRunner this is the main harness for the PizzaBot
func BotRunner(config cli.Config, path *grid.Path) {

	algorithm := algorithms.Algorithms[config.Algorithm]

	t := time.Now()
	paths := algorithm(path)
	duration := time.Since(t)

	results := make(map[int][]*grid.Path, 0)
	for _, p := range paths {
		results[len(p.Points)] = append(results[len(p.Points)], p)
	}

	if config.Verbose {
		WriteVerboseOutput(os.Stdout, results, duration)
	} else {
		WriteBasicOutput(os.Stdout, results)
	}
}
