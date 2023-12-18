package events

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

type ReactionStats struct {
	UserID    string
	Roles     []string
	Reactions map[string]int
}

var reactionStats = make(map[string]*ReactionStats)
var totalReactions int

func ReactionCounter(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	stats, ok := reactionStats[r.UserID]
	if !ok {
		stats = &ReactionStats{
			UserID:    r.UserID,
			Reactions: make(map[string]int),
		}
		reactionStats[r.UserID] = stats
	}

	stats.Reactions[r.Emoji.Name]++
	totalReactions++

	member, err := s.GuildMember(r.GuildID, r.UserID)
	if err != nil {
		log.Println("Erro ao obter membro do servidor:", err)
		return
	}
	stats.Roles = make([]string, len(member.Roles))
	for i, roleID := range member.Roles {
		role, err := s.State.Role(r.GuildID, roleID)
		if err != nil {
			log.Println("Erro ao obter cargo:", err)
			continue
		}
		stats.Roles[i] = role.Name
	}

	var reactions string
	for emojiName, reactionCount := range stats.Reactions {
		reactions += fmt.Sprintf("  - %s: %d reação(ões)\n", emojiName, reactionCount)
	}

	_, err = s.ChannelMessageSend("1101510837555974176", fmt.Sprintf("## Contador de Reações\n**Membro:** <@%s>,\n- Cargos:\n%s\n- Reações:\n%s\n- Total de reações: %d", stats.UserID, stats.Roles, reactions, totalReactions))
	if err != nil {
		log.Println("Erro ao enviar mensagem para o canal:", err)
	}
}
