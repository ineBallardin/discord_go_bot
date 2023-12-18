package events

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

type UserStats struct {
	UserID   string
	Roles    []string
	Messages map[string]int
}

var userStats = make(map[string]*UserStats)

func MessageCounter(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	stats, ok := userStats[m.Author.ID]
	if !ok {
		stats = &UserStats{
			UserID:   m.Author.ID,
			Messages: make(map[string]int),
		}
		userStats[m.Author.ID] = stats
	}

	stats.Messages[m.ChannelID]++

	member, err := s.GuildMember(m.GuildID, m.Author.ID)
	if err != nil {
		log.Println("Erro ao obter membro do servidor:", err)
		return
	}
	stats.Roles = make([]string, len(member.Roles))
	for i, roleID := range member.Roles {
		role, err := s.State.Role(m.GuildID, roleID)
		if err != nil {
			log.Println("Erro ao obter cargo:", err)
			continue
		}
		stats.Roles[i] = role.Name
	}

	isImpulserPro := false
	for _, role := range stats.Roles {
		if role == "impulserPRO" {
			isImpulserPro = true
			break
		}
	}

	var channels string
	for channelID, messageCount := range stats.Messages {
		channels += fmt.Sprintf("  - <#%s>: %d mensagem(s)\n", channelID, messageCount)
	}

	_, err = s.ChannelMessageSend("1101510837555974176", fmt.Sprintf("## Contador de Mensagens\n**Membro:** <@%s>,\n- **impulserPRO:** %v\n- Canais:\n%s", stats.UserID, isImpulserPro, channels))
	if err != nil {
		log.Println("Erro ao enviar mensagem para o canal:", err)
	}
}
