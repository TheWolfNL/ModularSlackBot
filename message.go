package bot

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/nlopes/slack"
)

type Message struct {
	*slack.MessageEvent
	bot         *Bot
	isBot       checkFunc
	isIM        checkFunc
	isMentioned checkFunc
}

type checkFunc func(*Message) bool

type MessageMock interface {
	GetUser() *slack.User
	GetChannel() *slack.Channel
	OpenIMChannel() string
}

func (bot *Bot) newMessage(event *slack.MessageEvent) *Message {
	return &Message{
		event,
		bot,
		isBot,
		isIM,
		isMentioned,
	}
}

func isBot(message *Message) bool {
	return message.User == "USLACKBOT" || message.BotID != ""
}

func (message *Message) IsBot() bool {
	return message.isBot(message)
}

func isIM(message *Message) bool {
	return message.Channel[0:1] == "D"
}

func (message *Message) IsIM() bool {
	return message.isIM(message)
}

func isMentioned(message *Message) bool {
	return strings.Contains(message.Text, "<@"+message.bot.id+">")
}

func (message *Message) IsMentioned() bool {
	return message.isMentioned(message)
}

func (message *Message) ToJson() string {
	messageJson, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		return "Invalid json"
	}
	return string(messageJson)
}

func (message *Message) GetUser() *slack.User {
	if message.bot == nil {
		return nil
	}
	if message.bot.slackApi == nil {
		return nil
	}
	user, error := message.bot.slackApi.GetUserInfo(message.User)
	if error != nil {
		log.Print(error)
	}
	return user
}

func (message *Message) GetChannel() *slack.Channel {
	if message.bot.slackApi == nil {
		return nil
	}
	channel, error := message.bot.slackApi.GetChannelInfo(message.Channel)
	if error != nil {
		log.Print(error)
	}
	return channel
}

func (message *Message) Respond(messageString string) {
	message.bot.SendMessage(message.Channel, messageString)
}

func (message *Message) OpenIMChannel(user string) string {
	return message.bot.OpenIMChannel(user)
}

func (bot *Bot) MockMessage(event *slack.MessageEvent, isBot checkFunc, isIM checkFunc, isMentioned checkFunc) *Message {
	return &Message{
		event,
		bot,
		isBot,
		isIM,
		isMentioned,
	}
}
