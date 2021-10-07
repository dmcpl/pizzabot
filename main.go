package main

import (
	"github.com/dmcpl/pizzabot/cli"
	"github.com/dmcpl/pizzabot/pizzabot"
	"log"
)

func main() {
	config, botParams, err := cli.ReadCommandLineArgs()
	if err != nil {
		log.Fatalf("Error reading command line args: %s", err)
	}

	path, err := cli.ParseBotParam(botParams)
	if err != nil {
		log.Fatalf("Experienced an error parsing Bot parameters: %s", err)
	}

	pizzabot.BotRunner(config, path)
}
