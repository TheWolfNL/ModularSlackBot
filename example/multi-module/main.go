package main

import (
	"fmt"
	"log"
	"os"

	"github.com/thewolfnl/ModularSlackBot"

	// Be sure to include the modules you want to use
	"github.com/TheWolfNL/ModularSlackBot-example-module-hello"
	"github.com/TheWolfNL/ModularSlackBot-example-module-reminder"
)

func main() {
	bot := bot.New("xoxb-123-ABC123")
	bot.AddModule(example.New())
	bot.AddModule(reminder.New())

	logger := log.New(os.Stdout, "messages-bot: ", log.Lshortfile|log.LstdFlags)
	bot.SetLogger(logger)

	// If you set debugging, it will log all requests to the console
	// Useful when encountering issues
	// bot.SetDebug(true)
	for _, channel := range bot.GetChannels() {
		fmt.Printf("ID: %s, Name: %s\n", channel.ID, channel.Name)
	}

	bot.Start()

}
