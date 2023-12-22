package events

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func MessageCounter(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Println("Contando...")
	if m.Author.Bot {
		return
	}

	member, err := s.GuildMember(m.GuildID, m.Author.ID)
	if err != nil {
		log.Println("Erro ao obter membro do servidor:", err)
		return
	}

	memberInfo := HandleEvent(s, member, m.GuildID)
	memberInfo.Messages++
	memberInfo.InfoChanged = true

	log.Println("TotalMessages for member", memberInfo.UserID, "is now", memberInfo.Messages)
}
