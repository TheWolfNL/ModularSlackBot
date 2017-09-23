package bot

type configuration struct {
	name     string
	version  string
	help     string
	triggers []Trigger
}

func NewModule(name string, version string) *Module {
	return &Module{
		config: configuration{
			name:    name,
			version: version,
			help: `
			{{.Bot.Name}} Module [{{.Bot.Version}}]
			There is no help text
			`,
		},
	}
}
