package test

import (
	"github.com/thewolfnl/ModularSlackBot/bot"
)

var module bot.Module

// New function to return a new instance of this bot
func New() bot.Module {
	module := bot.NewModule("Testbot", "0.0.1")

	// Define triggers
	module.AddTrigger("(?i)(hello|hi|hey).*", hello)

	module.AddTrigger("test", func(input string) {
		module.Respond("Read you loud and clear..")
	})

	return module
}

func hello(input string) {
	module.Respond("Hey there, to you too")
}
