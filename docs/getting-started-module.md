# Getting started making your own module
Making your own ModularSlackBot module is within your reach, no matter your technical expertise.
If your not new to GO and have it installed already, follow the steps at [Where to start](#where-to-start).

# Entirely new to GO?
Follow the instructions on [this page](https://golang.org/doc/install) to install GO.
I'd personally advise the use of [VS Code](https://code.visualstudio.com/) with the [GO plugin](https://marketplace.visualstudio.com/items?itemName=lukehoban.Go)

# Where to start
1. Fork the [hello example module](https://github.com/TheWolfNL/ModularSlackBot-example-module-hello) and clone it to your workspace
1. Install dependencies using [Glide](https://github.com/Masterminds/glide#install) by running `glide install`
1. First change the package name in `example.go` and `example_test.go`
1. Change the file names to match the package name
1. Add [triggers](triggers.md)
1. You can add tests by copying and altering `func ExampleHello()` in the `example_test.go` file.
1. To run tests call `go test -v` in commandline

## Public Module?
If you've created a public module you can get it listed on [the module page](modules.md).