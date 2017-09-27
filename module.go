package bot

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"regexp"

	"github.com/nlopes/slack"
)

type ModuleInterface interface {
	Name() string
	Version() string
	Status() bool
	AcceptsBotMessages() bool
	Help() string
	Info()
	HandleInput(*Message)
	Respond(string)
	AddTrigger(string, TriggerFunc)
	AcceptBotMessages()
	SetSlackApi(*slack.Client)
}

type Module struct {
	// Internal
	slack    *slack.Client
	triggers []Trigger
	channel  string

	// Properties
	name    string
	version string
	help    string

	// Options
	active            bool
	acceptBotMessages bool
}

func NewModule(name string, version string) *Module {
	return &Module{
		name:    name,
		version: version,
		help: `
		{{.Module.Name}} Module [{{.Module.Version}}]
		There is no help text
		`,
		active:            true,
		acceptBotMessages: false,
	}
}

func (module *Module) AcceptBotMessages() {
	module.acceptBotMessages = true
}

func (module *Module) AcceptsBotMessages() bool {
	return module.acceptBotMessages
}

func (module *Module) SetSlackApi(client *slack.Client) {
	module.slack = client
}

func (module *Module) AddTrigger(regex string, handler TriggerFunc) {
	trigger := Trigger{regex, func(message *Message) error {
		h := handler
		return h(message)
	}}
	module.triggers = append(module.triggers, trigger)
}

func (module *Module) HandleInput(message *Message) {
	trigger, err := module.checkForTriggers(message)
	if err == nil {
		module.SetChannel(message.Channel)
		trigger.handler(message)
	}
}

func (module *Module) SetChannel(channelId string) {
	module.channel = channelId
}

func (module *Module) checkForTriggers(message *Message) (*Trigger, error) {
	for _, trigger := range module.triggers {
		match, _ := regexp.MatchString(trigger.regex, message.Text)
		if match {
			return &trigger, nil
		}
	}
	return &Trigger{}, errors.New("No match")
}

func (module *Module) Help() string {
	// params := slack.PostMessageParameters{
	// 	AsUser: true,
	// }

	t := template.Must(template.New("help").Parse(module.help))
	var helpText bytes.Buffer
	data := struct {
		Module *Module
	}{
		Module: module,
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

func (module *Module) Status() bool {
	return module.active
}

func (module *Module) Name() string {
	return module.name
}

func (module *Module) Version() string {
	return module.version
}

func (module *Module) SetVersion(version string) {
	module.version = version
}

func (module *Module) Respond(message string) {
	if module == nil {
		return
	}
	fmt.Printf("\nSending to #%s\n", module.channel)
	fmt.Printf("Response: %s\n", message)
	if module.slack == nil {
		fmt.Print("Message not sent to slack because slack api is not configured\n")
	} else {
		params := slack.PostMessageParameters{
			AsUser: true,
		}

		channelID, timestamp, err := module.slack.PostMessage(module.channel, message, params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)
	}
}

func (module *Module) Info() {
	fmt.Println("Module Name: ", module.Name())
	fmt.Println("Module Version: ", module.Version())
	fmt.Println("Module Status: ", module.Status())
}
