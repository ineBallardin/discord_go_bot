package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type SlashCommand struct {
	Name        string
	Description string
	Handler     func(s *discordgo.Session, i *discordgo.InteractionCreate)
	Options     []*discordgo.ApplicationCommandOption
}

var SlashCommands []SlashCommand

func RegisterSlashCommands(s *discordgo.Session, guildID string) {
	for _, command := range SlashCommands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, guildID, &discordgo.ApplicationCommand{
			Name:        command.Name,
			Description: command.Description,
			Options:     command.Options,
		})
		if err != nil {
			log.Printf("Cannot create '%v' command: %v", command.Name, err)
		} else {
			log.Printf("Command '%v' registered successfully", command.Name)
		}
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionApplicationCommand {
			for _, command := range SlashCommands {
				if i.ApplicationCommandData().Name == command.Name {
					command.Handler(s, i)
					break
				}
			}
		}
	})
}
