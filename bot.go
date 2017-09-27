package bot

import (
	"fmt"
	"log"

	"github.com/nlopes/slack"
)

type Bot struct {
	slackApi *slack.Client
	modules  []ModuleInterface
}

func New(slackBotToken string) *Bot {
	return &Bot{
		slackApi: slack.New(slackBotToken),
	}
}

func (bot *Bot) AddModule(module ModuleInterface) {
	module.SetSlackApi(bot.slackApi)
	bot.modules = append(bot.modules, module)
}

func (bot *Bot) SetDebug(value bool) {
	bot.slackApi.SetDebug(value)
}

func (bot *Bot) SetLogger(logger *log.Logger) {
	slack.SetLogger(logger)
}

func (bot *Bot) GetChannels() []slack.Channel {
	channels, err := bot.slackApi.GetChannels(false)
	if err != nil {
		fmt.Printf("%s\n", err)
		return []slack.Channel{}
	}
	return channels
}

func (bot *Bot) Start() {
	rtm := bot.slackApi.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
			case *slack.TeamJoinEvent:
				// Handle new user to client
			case *slack.MessageEvent: //
				// Handle new message to channel
				message := Message{ev}
				// fmt.Printf("Message: %s \n", message.ToJson())

				for _, module := range bot.modules {
					if (module.AcceptsBotMessages() && message.IsBot()) || !message.IsBot() {
						module.HandleInput(&message)
					}
				}
			case *slack.ReactionAddedEvent:
				// Handle reaction added
			case *slack.ReactionRemovedEvent:
				// Handle reaction removed
			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())
			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop
			default:
				// fmt.Printf("Unknown error")
			}
		}
	}
}
