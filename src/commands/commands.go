package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type SlashCommand struct {
	Name        string
	Description string
	Handler     func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

var SlashCommands []SlashCommand

func RegisterSlashCommands(s *discordgo.Session, guildID string) {
	for _, command := range SlashCommands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, guildID, &discordgo.ApplicationCommand{
			Name:        command.Name,
			Description: command.Description,
		})
		if err != nil {
			log.Printf("Cannot create '%v' command: %v", command.Name, err)
		}

		s.AddHandler(command.Handler)
	}
}
