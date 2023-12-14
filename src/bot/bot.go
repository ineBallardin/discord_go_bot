package bot

type Command struct {
	Data    string
	Execute func()
}

type Bot struct {
	Cooldowns map[string]string
	Commands  map[string]Command
}

func createBot() *Bot {
	return &Bot{
		Cooldowns: make(map[string]string),
		Commands:  make(map[string]Command),
	}
}
