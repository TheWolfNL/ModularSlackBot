package bot

type Trigger struct {
	regex   string
	handler func(message Message)
}
