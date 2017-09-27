package bot

type Trigger struct {
	regex   string
	handler TriggerFunc
}

type TriggerFunc func(message *Message) error
