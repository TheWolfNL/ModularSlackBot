package main

import (
	"fmt"

	// Be sure to your own module here
	"github.com/TheWolfNL/ModularSlackBot-example-module-hello"
	"github.com/thewolfnl/ModularSlackBot"
)

var questions = []string{
	"Test 1, 2, 3.",
	"test is case sensitive",
	"Hello bot",
	"Hey bot",
	"random stuff",
}

func main() {
	module := example.New()
	module.Info()

	fmt.Print("\nTriggering sample questions:\n")
	for _, question := range questions {
		fmt.Printf("\nYour question was: '%s'\n", question)
		module.HandleInput(bot.CreateMessage(question))
	}

}
