package botTestLib

import (
	"encoding/json"

	"github.com/nlopes/slack"
	"github.com/thewolfnl/ModularSlackBot"
)

type MockBot struct {
	*bot.Bot
}

func NewMockBot() *MockBot {
	return &MockBot{&bot.Bot{}}
}

type checkFunc func(*bot.Message) bool

func ReturnFalse(_ *bot.Message) bool {
	return false
}

func ReturnTrue(_ *bot.Message) bool {
	return true
}

func (mock *MockBot) CreateMessage(messageString string) *bot.Message {
	return mock.MockMessage(mock.CreateMessageEvent(messageString), ReturnFalse, ReturnFalse, ReturnFalse)
}

func (mock *MockBot) CreateMessageEvent(messageString string) *slack.MessageEvent {
	message := &slack.MessageEvent{}
	messageJson := `{
		"type":"message",
		"channel": "C2147483705",
		"user": "U2147483697",
		"text": "Hello world",
		"ts":"1355517523.000005"
		}`
	if err := json.Unmarshal([]byte(messageJson), &message); err != nil {
		return &slack.MessageEvent{}
	}
	message.Text = messageString
	return message
}
