# ModularSlackBot
This wil be a basic slack bot that allows you to easily add modules containing responses

Before use:
`glide install` to install dependencies using glide

Set the bot token in `main.go`

`go run main.go` to start this bot

make changes to `modules/test.go`

# Roadmap
v1.0
- move RTM to bot
- improve project structure
- create tests
- create example project
- Solely Trigger based

v1.1
- allow for more interaction by giving access to message event (to access username)
- add time based triggers (cron)

v1.2
- add interactive mode