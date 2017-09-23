package reminder

import (
	"fmt"

	"github.com/thewolfnl/ModularSlackBot/bot"
)

var reminders []reminder

// New function to return a new instance of this bot
func New() ReminderModule {
	reminder := ReminderModule{bot.NewModule("Reminder", "0.0.1")}

	// Define triggers
	reminder.AddTrigger("(?i)notify #[0-9]+.*", reminder.pipeline)
	reminder.AddTrigger("(?i)remind #[0-9]+.*", reminder.pipelineReminder)

	return reminder
}

func (module *ReminderModule) pipelineReminder(message bot.Message) {
	if !message.IsBot() {
		reminders = append(reminders, reminder{1234, message.User})
		module.SetChannel(message.Channel)
		module.Respond("Seting reminder for '" + message.Text + "' " + message.Username)
		fmt.Print(reminders)
	}
}

func (module *ReminderModule) pipeline(message bot.Message) {
	if message.IsBot() {
		fmt.Printf("\nMessageJSON: \n%s\n", message.ToJson())
		module.SetChannel(message.Channel)
		module.Respond("Notification from reminder")
	}
}
