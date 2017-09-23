# Roadmap
v1.0
- √ improve project structure
- √ create example module project
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
- create bot example
- create tests
- Solely Trigger based

v1.1
- add time based triggers (cron)

v1.2
- add interactive mode
