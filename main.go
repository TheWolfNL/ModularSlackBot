package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/nlopes/slack"
	"github.com/thewolfnl/ModularSlackBot/modules"
)

var questions = []string{
	"Test 1, 2, 3.",
	"Hello bot",
	"Hey bot",
	"random stuff",
}

func main() {
	api := slack.New("bot-token")
	test := test.New()
	test.SetSlackApi(api)

	logger := log.New(os.Stdout, "messages-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)

	// If you set debugging, it will log all requests to the console
	// Useful when encountering issues
	// api.SetDebug(true)
	channels, err := api.GetChannels(false)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, channel := range channels {
		fmt.Printf("ID: %s, Name: %s\n", channel.ID, channel.Name)
	}

	rtm := api.NewRTM()
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

				// Only respond to real users. Bots have BotIDs, users do not
				if ev.Msg.BotID == "" {

					message, err := json.Marshal(ev.Msg)
					if err != nil {
						fmt.Println(err)
						return
					}
					fmt.Printf("Message: %s\nJson: %s \n", ev.Msg, string(message))
					test.SetChannel(ev.Msg.Channel)
					test.HandleInput(ev.Msg.Text)
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

	// test.Info()
	// fmt.Printf("Bot Help: \n%s", test.Help("channelid 1"))

	// fmt.Print("\nTriggering sample questions:\n")
	// for _, question := range questions {
	// 	fmt.Printf("\nYour question was: '%s'\n", question)
	// 	test.HandleInput(question)
	// }

}
