package cli

import (
	"example.com/slice/pizzabot/algorithms"
	"flag"
	"fmt"
	"strings"
)

type Config struct {
	Verbose      bool
	Algorithm    string
}

func ReadCommandLineParams() Config {
	var config Config

	flag.BoolVar(&config.Verbose, "verbose", false, "Verbose output (include all paths attempted). Default off.")
	flag.BoolVar(&config.Verbose, "v", false, "Verbose output (include all paths attempted). Default off.")

	flag.StringVar(&config.Algorithm, "algorithm", "CP", fmt.Sprintf("Which algorithm to chose from one of %s", strings.Join(algorithms.AvailableAlgorithms(), ", ")))
	flag.StringVar(&config.Algorithm, "a", "CP", fmt.Sprintf("Which algorithm to chose from one of %s", strings.Join(algorithms.AvailableAlgorithms(), ", ")))
	//todo validate this input


	flag.Parse()

	return config
}