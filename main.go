package main

import (
	"example.com/slice/pizzabot/cli"
	"example.com/slice/pizzabot/pizzabot"
	"log"
	"os"
)

func main() {
	config := cli.ReadCommandLineParams()
	path, err := cli.ParseBotParam(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatalf("Experienced an error parsing Bot parameters: %s", err)
	}

	pizzabot.BotRunner(config, path)
}
