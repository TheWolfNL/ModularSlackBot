# Triggers
Use the following code and change the repsonse message and trigger [regular expression](#regular-expressions).
```
module.AddTrigger("ping", func(message bot.Message) {
    module.Respond("pong")
})
```
This example trigger will listen to "ping" inside a message. You can use regular expressions to define this trigger.

## Features
It's very likely that during development you'd want additional functionality, don't hesitate to create an [issue](https://github.com/TheWolfNL/ModularSlackBot/issues/new) (please add the `feature request` label), and I'll try to help you out.

## Problems
If it's your first GO project is can de a bit discouraging because the strictness will give errors, if you google the error you'll find some answers.
Please keep in mind that when you're googling for a solution you sometimes need to substitute `go` for `golang` to avoid undesired results. 
If google doesn't solve your problem, please create an [issue](https://github.com/TheWolfNL/ModularSlackBot/issues/new) and I'll try to help.

## Regular Expressions
Use a website like [Regex101](https://regex101.com/) and select `Golang` as flavor.

Copy the following lines into the `test string` area
```
Asswooping
ping
typing
ping google.com
PingPong
PING
Sentences containing ping?
```

If you then type `ping` in the regular expression field, You'll see 5 highlights, these sentences will trigger the function.

* You can also notice it is case sensitive, change the expression to `(?i)ping` to solve this.
* To make sure `ping` is a word on its own we can change it to `\bping\b`

    This `\b` defines a word boundary.
    * It can also be used to find words ending in `ping` by using `\bping`.
    * Or words starting with `ping` by using `ping\b`.
    * Combined with case sensitivity it will become `(?i)\bping\b`
* You're also able to have multiple keywords `(\bping\b|Pong)`

    This will only trigger if `ping` is a word on it's own or `Pong` is found
* To make sure `ping` is the only thing sent we can use `^ping$` to check

    Make sure the `multi-line` flag is active by clicking on the flag at the end of the input field
    * It's also possible to only check for the start of the string `^ping`
    * Ot at the end of the string `ping$`

There is a lot more possible when using regular expressions, but this site will help you test it.