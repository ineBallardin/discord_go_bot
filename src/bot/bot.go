package bot

import "github.com/bwmarrin/discordgo"

type Command struct {
	Data    string
	Execute func()
}

type Bot struct {
	Session   *discordgo.Session
	Cooldowns map[string]string
	Commands  map[string]Command
}

func CreateBot() *Bot {
	return &Bot{
		Cooldowns: make(map[string]string),
		Commands:  make(map[string]Command),
	}
}
