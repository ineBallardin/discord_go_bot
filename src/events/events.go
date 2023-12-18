package events

import "github.com/bwmarrin/discordgo"

func RegisterEvents(s *discordgo.Session) {
	s.AddHandler(PingMessageCreate)
	s.AddHandler(MessageCounter)
	s.AddHandler(ReactionCounter)
	s.AddHandler(VoiceChannelCounter)
	// Adicionar mais manipuladores de eventos aqui cfe necessário
}
