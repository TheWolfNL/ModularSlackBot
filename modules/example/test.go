package example

import (
	"fmt"

	"github.com/thewolfnl/ModularSlackBot/bot"
)

type TestModule struct {
	*bot.Module
}

// New function to return a new instance of this bot
func New() TestModule {
	module := TestModule{bot.NewModule("Testbot", "0.0.1")}

	// Define triggers
	module.AddTrigger("(?i)(hello|hi|hey).*", module.hello)

	module.AddTrigger("test", func(message bot.Message) {
		fmt.Printf("\nMessageJSON: \n%s\n", message.ToJson())
		module.Respond("Read you loud and clear..")
	})

	return module
}

func (module *TestModule) hello(message bot.Message) {
	module.Respond("Hey there, to you too")
}
