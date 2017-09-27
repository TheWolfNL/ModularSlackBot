package bot

import (
	"encoding/json"
	"fmt"

	"github.com/nlopes/slack"
)

type Message struct {
	*slack.MessageEvent
}

func (message *Message) IsBot() bool {
	return message.User == "USLACKBOT" || message.BotID != ""
}

func (message *Message) ToJson() string {
	messageJson, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		return "Invalid json"
	}
	return string(messageJson)
}

func CreateMessage(messageString string) *Message {
	message := &Message{}
	messageJson := `{
		"channel": "C2147483705",
		"user": "U2147483697",
		"text": "Hello world"
		}`
	if err := json.Unmarshal([]byte(messageJson), &message); err != nil {
		return &Message{}
	}
	message.Text = messageString
	return message
}
