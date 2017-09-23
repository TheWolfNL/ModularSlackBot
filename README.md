# ModularSlackBot
This wil be a basic slack bot that allows you to easily add modules containing responses

Before use:
`glide install` to install dependencies using glide

Set the bot token in `main.go`

`go run main.go` to start this bot

make changes to `modules/test.go`

# Roadmap
v1.0
- √ improve project structure
- change structs
    - Bot
        Add functions to allow module management on the bot
        - integrate RTM into bot
        - let bot handle 
        - activating/deactivating a module
        - triggering module info
        - triggering module help
        Contains interaction with slack, default functions to allow basic desired functionality

        - change Message struct to provide the actual channel name
        - change Message struct to provide the actual user name

    - Module
        - change Module struct to remove configuration struct because it should be part of the module itself
- create example bot project
- create example module project
- create tests
- Solely Trigger based

v1.1
- add time based triggers (cron)

v1.2
- add interactive mode
