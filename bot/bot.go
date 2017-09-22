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
	slack   *slack.Client
	config  configuration
	channel string
}

type configuration struct {
	name     string
	version  string
	help     string
	triggers []Trigger
}

type Trigger struct {
	regex   string
	handler func(string)
}

func NewModule(name string, version string) Module {
	return Module{
		config: configuration{
			name:    name,
			version: version,
			help: `
			{{.Bot.Name}} Module [{{.Bot.Version}}]
			There is no help text
			`,
		},
	}
}

var slackClient *slack.Client

func (bot *Module) SetSlackApi(client *slack.Client) {
	bot.slack = client
	slackClient = client
}

var channel string

func (bot *Module) SetChannel(channelId string) {
	bot.channel = channelId
	channel = channelId
	fmt.Print("Channel", channel, bot.channel)
}

func (bot *Module) AddTrigger(regex string, handler func(string)) {
	trigger := Trigger{regex, handler}
	bot.config.triggers = append(bot.config.triggers, trigger)
}

func (bot *Module) HandleInput(input string) {
	trigger, err := bot.checkForTriggers(input)
	if err == nil {
		trigger.handler(input)
	}
}

func (bot *Module) checkForTriggers(input string) (Trigger, error) {
	for _, trigger := range bot.config.triggers {
		match, _ := regexp.MatchString(trigger.regex, strings.ToLower(input))
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
	fmt.Print("Channel", channel, bot.channel)
	fmt.Printf("\nSending '%s' to #%s\n", message, channel)
	if slackClient == nil {
		fmt.Printf("Response: %s\n", message)
	} else {
		params := slack.PostMessageParameters{
			AsUser: true,
		}

		channelID, timestamp, err := slackClient.PostMessage(channel, message, params)
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
