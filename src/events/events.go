package events

import "github.com/bwmarrin/discordgo"

func RegisterEvents(s *discordgo.Session) {
	s.AddHandler(MessageCreate)
	// Adicionar mais manipuladores de eventos aqui cfe necessário
}
