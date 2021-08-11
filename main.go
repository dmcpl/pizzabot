package main

import (
	"example.com/slice/pizzabot/cli"
	"example.com/slice/pizzabot/pizzabot"
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
