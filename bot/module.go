package bot

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"regexp"
	"strings"

	"github.com/nlopes/slack"
)

type Module struct {
	Slack   *slack.Client
	config  configuration
	Channel string
}

func (bot *Module) SetSlackApi(client *slack.Client) {
	bot.Slack = client
}

func (bot *Module) AddTrigger(regex string, handler func(message Message)) {
	trigger := Trigger{regex, handler}
	bot.config.triggers = append(bot.config.triggers, trigger)
}

func (bot *Module) HandleInput(message Message) {
	trigger, err := bot.checkForTriggers(message)
	if err == nil {
		trigger.handler(message)
	}
}

func (bot *Module) SetChannel(channelId string) {
	bot.Channel = channelId
}

func (bot *Module) checkForTriggers(message Message) (Trigger, error) {
	for _, trigger := range bot.config.triggers {
		match, _ := regexp.MatchString(trigger.regex, strings.ToLower(message.Text))
		if match {
			return trigger, nil
		}
	}
	return Trigger{}, errors.New("No match")
}

func (bot *Module) Help(channelId string) string {
	// params := slack.PostMessageParameters{
	// 	AsUser: true,
	// }

	t := template.Must(template.New("help").Parse(bot.config.help))
	var helpText bytes.Buffer
	data := struct {
		Bot     *Module
		Channel string
	}{
		Bot:     bot,
		Channel: channelId,
	}
	if err := t.Execute(&helpText, data); err != nil {
		return err.Error()
	}

	return helpText.String()

	// channelID, timestamp, err := bot.slack.PostMessage(channelId, helpText, params)
	// if err != nil {
	// 	fmt.Printf("%s\n", err)
	// 	return
	// }
	// fmt.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)
}

func (bot *Module) Status() string {
	return "Active"
}

func (bot *Module) Name() string {
	return bot.config.name
}

func (bot *Module) Version() string {
	return bot.config.version
}

func (bot *Module) SetVersion(version string) {
	bot.config.version = version
}

func (bot *Module) Respond(message string) {
	fmt.Printf("\nSending '%s' to #%s\n", message, bot.Channel)
	if bot.Slack == nil {
		fmt.Printf("Response: %s\n", message)
	} else {
		params := slack.PostMessageParameters{
			AsUser: true,
		}

		channelID, timestamp, err := bot.Slack.PostMessage(bot.Channel, message, params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)
	}
}

func (bot *Module) Info() {
	fmt.Println("Bot Name: ", bot.Name())
	fmt.Println("Bot Version: ", bot.Version())
	fmt.Println("Bot Status: ", bot.Status())
}
