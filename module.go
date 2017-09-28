package bot

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"regexp"
)

type Module struct {
	// Internal
	bot      *Bot
	triggers []Trigger

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

func (module *Module) SetBot(bot *Bot) {
	module.bot = bot
}

func (module *Module) AddTrigger(regex string, handler TriggerFunc) {
	trigger := Trigger{regex, func(message *Message) {
		h := handler
		h(message)
	}}
	module.triggers = append(module.triggers, trigger)
}

func (module *Module) HandleInput(message *Message) {
	trigger, err := module.checkForTriggers(message)
	if err == nil {
		trigger.handler(message)
	}
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

func (module *Module) Info() {
	fmt.Println("Module Name: ", module.Name())
	fmt.Println("Module Version: ", module.Version())
	fmt.Println("Module Status: ", module.Status())
}
