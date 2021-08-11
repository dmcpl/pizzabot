package cli

import (
	"example.com/slice/pizzabot/algorithms"
	"flag"
	"fmt"
	"strings"
)

type Config struct {
	Verbose   bool
	Algorithm string
}

// ReadCommandLineArgs read in command line switches and bot run parameters
func ReadCommandLineArgs() (config Config, botParams string, err error) {
	flag.BoolVar(&config.Verbose, "-verbose", false, "Verbose output (include all paths attempted). Default off.")
	flag.BoolVar(&config.Verbose, "v", false, "Verbose output (include all paths attempted). Default off.")

	flag.StringVar(&config.Algorithm, "-algorithm", "CP", fmt.Sprintf("Which algorithm to chose from one of: [%s]", strings.Join(algorithms.AvailableAlgorithms(), ", ")))
	flag.StringVar(&config.Algorithm, "a", "CP", fmt.Sprintf("Which algorithm to chose from one of: [%s]", strings.Join(algorithms.AvailableAlgorithms(), ", ")))

	flag.Parse()

	if botParams = flag.Arg(0); len(botParams) != 0 {
		return config, botParams, nil
	}

	return config, botParams, fmt.Errorf("no bot run parameters found")
}
