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
