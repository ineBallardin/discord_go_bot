package commands

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	SlashCommands = append(SlashCommands, SlashCommand{
		Name:        "ping",
		Description: "Responds with Pong!",
		Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.Type == discordgo.InteractionApplicationCommand {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Pong!",
					},
				})
			}
		},
	})
}
