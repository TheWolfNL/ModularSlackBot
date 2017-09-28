package bot

import (
	"fmt"
	"log"

	"github.com/nlopes/slack"
)

type Bot struct {
	slackApi *slack.Client
	id       string
	modules  []*Module
}

func New(slackBotToken string) *Bot {
	bot := &Bot{
		slackApi: slack.New(slackBotToken),
	}
	response, err := bot.slackApi.AuthTest()
	if err != nil {
		fmt.Print(err)
	} else {
		bot.id = response.UserID
	}
	return bot
}

func (bot *Bot) AddModule(module *Module) {
	module.SetBot(bot)
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
				message := bot.newMessage(ev)
				// fmt.Printf("Message: %s \n", message.ToJson())

				for _, module := range bot.modules {
					if (module.AcceptsBotMessages() && message.IsBot()) || !message.IsBot() {
						module.HandleInput(message)
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

func (bot *Bot) SendMessage(channel string, message string) {
	if bot == nil {
		return
	}
	fmt.Printf("\nSending to #%s\n", channel)
	fmt.Printf("Response: %s\n", message)
	if bot.slackApi == nil {
		fmt.Print("Message not sent to slack because slack api is not configured\n")
	} else {
		params := slack.PostMessageParameters{
			AsUser: true,
		}

		channelID, timestamp, err := bot.slackApi.PostMessage(channel, message, params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)
	}
}

func (bot *Bot) OpenIMChannel(user string) string {
	noOp, alreadyOpen, channel, error := bot.slackApi.OpenIMChannel(user)
	fmt.Print(noOp, alreadyOpen, channel, error)
	if error == nil {
		return channel
	}
	return ""
}
