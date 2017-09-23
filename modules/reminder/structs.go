package reminder

import "github.com/thewolfnl/ModularSlackBot/bot"

type ReminderModule struct {
	*bot.Module
}

type reminder struct {
	buildID int
	user    string
}
